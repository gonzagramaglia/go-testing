package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

/*
What have we done here?

We've refactored our assertion into a new function.
This reduces duplication and improves the readability of our tests.

We need to pass in t *testing.T so that we can tell the test code
to fail when we need to.

For helper functions, it's a good idea to accept a testing.TB
which is an interface that *testing.T and *testing.B both satisfy,
so you can call helper functions from a test, or a benchmark.

t.Helper() is needed to tell the test suite that this method is
a helper. By doing this, when it fails, the line number reported
will be in our function call rather than inside our test helper.
This will help other developers track down problems more easily.

When you have more than one argument of the same type
(in our case two strings) rather than having
(got string, want string) you can shorten it to (got, want string).
*/
