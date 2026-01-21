package greeting

import (
	"fmt"
	"unicode"
)

var Greet = greet

func greet(name ...string) string {
	if len(name) == 0 {
		return "Hello, my friend."
	}

	firstname := []rune(name[0])

	if unicode.IsUpper(firstname[len(firstname)-1]) {
		return fmt.Sprintf("HELLO, %s.", name[0])
	}

	return fmt.Sprintf("Hello, %s.", name[0])
}
