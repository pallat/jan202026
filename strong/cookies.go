package main

import (
	"fmt"
	"net/http"
)

// AuthMiddleware ทำหน้าที่ตรวจสอบ Cookie
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. ดึง Cookie ชื่อ "session_token" จาก Request
		cookie, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				// ถ้าไม่มี Cookie ส่งกลับ 401
				http.Error(w, "Unauthorized: No cookie found", http.StatusUnauthorized)
				return
			}
			// ถ้าเกิด Error อื่นๆ
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// 2. ตรวจสอบความถูกต้องของ Token (ในที่นี้สมมติว่าเช็คกับ Database หรือ Cache)
		sessionToken := cookie.Value
		if !isValidSession(sessionToken) {
			http.Error(w, "Unauthorized: Invalid session", http.StatusUnauthorized)
			return
		}

		// 3. ถ้าผ่านด่าน ให้ไปทำงานที่ Handler ถัดไป
		next.ServeHTTP(w, r)
	})
}

// ฟังก์ชันสมมติสำหรับเช็คความถูกต้องของ Token
func isValidSession(token string) bool {
	// ในงานจริง: เช็คใน Redis, Database หรือ Parse JWT
	return token == "random-secure-session-id"
}

// วิธีการนำไปใช้ใน Main
func main() {
	mux := http.NewServeMux()

	// API ที่ต้องการการตรวจสอบสิทธิ์
	finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ยินดีด้วย! คุณเข้าถึงข้อมูลลับของ สสช. ได้"))
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "index.html")
	})

	mux.HandleFunc("/login", LoginHandler)
	// ห่อหุ้ม Handler ด้วย Middleware
	mux.Handle("/v1/data", AuthMiddleware(finalHandler))

	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// 1. ตรวจสอบ Username/Password (สมมติว่าผ่าน)
	sessionToken := "random-secure-session-id"

	// 2. สร้าง Cookie พร้อมการป้องกันครบสูตร
	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Path:     "/",
		HttpOnly: true,                 // ป้องกัน JavaScript
		Secure:   false,                // ตั้งเป็น false สำหรับ local testing (HTTP)
		SameSite: http.SameSiteLaxMode, // ป้องกัน CSRF
		MaxAge:   86400,                // อายุ 1 วัน
	}

	// 3. ส่ง Cookie กลับไปที่ Browser
	http.SetCookie(w, cookie)
	fmt.Fprintln(w, "Login Successful")
}
