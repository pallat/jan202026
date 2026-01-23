package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	lvl := os.Getenv("LOG_LEVEL")
	switch lvl {
	case "debug":
		slog.SetLogLoggerLevel(slog.LevelDebug)
	case "info":
		slog.SetLogLoggerLevel(slog.LevelInfo)
	default:
		slog.SetLogLoggerLevel(slog.LevelError)
	}

	var err error
	db, err = sql.Open("sqlite3", "./oscar.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	handler := NewOscarHandler(db)
	r.POST("/oscar/insert", handler.Insert)

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/provinces/:provinces", proviceHandler)
	r.POST("/oscar/years", oscarOfTheYearHandler)
	r.POST("/oscars", oscarHandler)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

type OscarHandler struct {
	db *sql.DB
}

func NewOscarHandler(db *sql.DB) *OscarHandler {
	return &OscarHandler{db: db}
}

func (handler *OscarHandler) Insert(c *gin.Context) {
	var input maleOscar
	if err := c.Bind(&input); err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()
	if _, err := handler.db.ExecContext(ctx, "insert into reward (name, movie_name) values (?,?)", input.Name, input.MovieName); err != nil {
		slog.Error(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

}

type year struct {
	Year string `json:"year"`
}

func oscarOfTheYearHandler(ctx *gin.Context) {
	var y year
	if err := ctx.Bind(&y); err != nil {
		slog.Error(err.Error())
		return
	}

	f, err := os.Open("oscar.csv")
	if err != nil {
		slog.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		slog.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	for _, record := range records {
		if record[1] == y.Year {
			age, err := strconv.Atoi(record[2])
			if err != nil {
				slog.Error(err.Error())
			}
			ctx.JSON(http.StatusOK, maleOscar{
				Year:      y.Year,
				Age:       age,
				Name:      record[3],
				MovieName: record[4],
			})
		}
	}

	ctx.Status(http.StatusNotFound)
}

type maleOscar struct {
	Year      string `json:"year"`
	Age       int    `json:"age"`
	Name      string `json:"name"`
	MovieName string `json:"movie_name"`
}

func oscarHandler(ctx *gin.Context) {
	var mo maleOscar
	if err := ctx.Bind(&mo); err != nil {
		slog.Error(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, mo)
}

func proviceHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response{
		Code:    "20010",
		Message: ctx.Param("provinces"),
	})
}

type response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
