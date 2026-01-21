package foobar

import "fmt"

func FizzBuzz(n int) string {
	if n%3 == 0 {
		return "Fizz"
	} else if n == 5 {
		return "Buzz"
	}
	return fmt.Sprint(n)
}
