package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("to a person", func(t *testing.T) {
		got := Hello("Vinoop", "")
		want := "Hello, Vinoop"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Vinoop", "Spanish")
		want := "Hola, Vinoop"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Vinoop", "French")
		want := "Bonjour, Vinoop"
		assertCorrectMessage(t, got, want)
	})

}
