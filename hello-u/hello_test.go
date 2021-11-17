package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Juan")
	want := "Hello, Juan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
