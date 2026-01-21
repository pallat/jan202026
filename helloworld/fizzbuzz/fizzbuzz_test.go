package fizzbuzz_test

import (
	"helloworld/fizzbuzz"
	"testing"
)

func TestFizzBuzzOne(t *testing.T) {
	given := 1
	want := "1"

	actual := fizzbuzz.FizzBuzz(given)

	if want != actual {
		t.Errorf("it should be %q but %q\n", want, actual)
	}
}
func TestFizzBuzzTwo(t *testing.T) {
	given := 2
	want := "2"

	actual := fizzbuzz.FizzBuzz(given)

	if want != actual {
		t.Errorf("it should be %q but %q\n", want, actual)
	}
}
func TestFizzBuzzThree(t *testing.T) {
	given := 3
	want := "Fizz"

	actual := fizzbuzz.FizzBuzz(given)

	if want != actual {
		t.Errorf("it should be %q but %q\n", want, actual)
	}
}
func TestFizzBuzzSix(t *testing.T) {
	given := 6
	want := "Fizz"

	actual := fizzbuzz.FizzBuzz(given)

	if want != actual {
		t.Errorf("it should be %q but %q\n", want, actual)
	}
}
func TestFizzBuzzNine(t *testing.T) {
	given := 9
	want := "Fizz"

	actual := fizzbuzz.FizzBuzz(given)

	if want != actual {
		t.Errorf("it should be %q but %q\n", want, actual)
	}
}
func TestFizzBuzzFive(t *testing.T) {
	given := 5
	want := "Buzz"

	actual := fizzbuzz.FizzBuzz(given)

	if want != actual {
		t.Errorf("it should be %q but %q\n", want, actual)
	}
}

// 1 -> "1"
// 2 -> "2"
// 3 -> "Fizz"
// 4 -> "4"
// 5 -> "Buzz"
