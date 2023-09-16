package main

import "testing"

// File name should be xxx_test.go
// Test func should start with "Test"
// Test func takes only one argument "t *testing.T"
// Need to import "testing"
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)

	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Mihael", "French")
		want := "Bonjour, Mihael"
		assertCorrectMessage(t, got, want)
	})
}

// For helper functions, it's a good idea to accept a testing.TB which is an interface that *testing.T and *testing.B both satisfy, so you can call helper functions from a test, or a benchmar
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
