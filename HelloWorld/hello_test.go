package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		t.Run("in english", func(t *testing.T) {
			got := Hello("Chris", "English")
			want := "Hello, Chris"
			assertCorrectMessage(t, got, want)
		})
		t.Run("in Spanish", func(t *testing.T) {
			got := Hello("Elodie", "Spanish")
			want := "Hola, Elodie"
			assertCorrectMessage(t, got, want)
		})
		t.Run("in French", func(t *testing.T) {
			got := Hello("Emilie", "French")
			want := "Bonjour, Emilie"
			assertCorrectMessage(t, got, want)
		})
	})
	t.Run("saying hello to world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper() // This tells the testing package that this function is a helper function
	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
