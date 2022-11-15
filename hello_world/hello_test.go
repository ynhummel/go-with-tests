package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Yuri", "")
		want := "Hello, Yuri"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Yuri", "French")
		want := "Bonjour, Yuri"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("Yuri", "Portuguese")
		want := "Olá, Yuri"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
