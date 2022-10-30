package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to person in English", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to person in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to person in French", func(t *testing.T) {
		got := Hello("Lauren", "French")
		want := "Bonjour, Lauren"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to world in English when name is an empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
