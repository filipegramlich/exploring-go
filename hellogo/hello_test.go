package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Filipe", "")
		want := "Hello Filipe!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people when and empty string	is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello GO!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Português", func(t *testing.T) {
		got := Hello("Filipe", "Português")
		want := "Olá Filipe!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("French", func(t *testing.T) {
		got := Hello("Filipe", "French")
		want := "Bonjour Filipe!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
