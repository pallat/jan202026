package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()

	wg.Add(5)
	for i := range 5 {
		go slowPrint(i)
	}

	wg.Wait()
	fmt.Println(time.Since(start))
}

func slowPrint(ch chan int, i int) {
	ch <- 10
	x := <-ch
	time.Sleep(100 * time.Millisecond)
	fmt.Println(i)
	wg.Done()
}
