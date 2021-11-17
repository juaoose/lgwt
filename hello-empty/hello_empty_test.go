package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		// This marks the function as a helper, so that if the test fails the caller is reported instead of this function
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Juan")
		want := "Hello, Juan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

}
