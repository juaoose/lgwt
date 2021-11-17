package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to Juan in spanish returns 'Hola, Juan'", func(t *testing.T) {
		got := Hello("Juan", "Spanish")
		want := "Hola, Juan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to Jason in french returns 'Bonjour, Jason'", func(t *testing.T) {
		got := Hello("Jason", "French")
		want := "Bonjour, Jason"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to James in english returns 'Hello, James'", func(t *testing.T) {
		got := Hello("James", "English")
		want := "Hello, James"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello supplying empty parameters returns 'Hello, World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello supplying an empty name and a spanish as language 'Hola, World'", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello supplying to Jenny returns 'Hello, Jenny'", func(t *testing.T) {
		got := Hello("Jenny", "")
		want := "Hello, Jenny"
		assertCorrectMessage(t, got, want)
	})

}
