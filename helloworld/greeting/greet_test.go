package greeting

import (
	"testing"
)

func TestGreetAName(t *testing.T) {
	given := "Bob"
	want := "Hello, Bob."

	actual := greet(given)

	if want != actual {
		t.Errorf("it should be %q but actual %q\n", want, actual)
	}
}

func TestGreetAnotherName(t *testing.T) {
	given := "Jane"
	want := "Hello, Jane."

	actual := greet(given)

	if want != actual {
		t.Errorf("it should be %q but actual %q\n", want, actual)
	}
}

func TestGreetMyFriend(t *testing.T) {
	want := "Hello, my friend."

	actual := greet()

	if want != actual {
		t.Errorf("it should be %q but actual %q\n", want, actual)
	}
}

func TestShoutOutTheName(t *testing.T) {
	given := "JERRY"
	want := "HELLO, JERRY."

	actual := greet(given)

	if want != actual {
		t.Errorf("it should be %q but actual %q\n", want, actual)
	}
}
