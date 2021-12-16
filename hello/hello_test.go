package main

import "testing"

func TestHello(t *testing.T) {
	verifyMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Caio")
		want := "Hello, Caio"

		verifyMessage(t, got, want)
	})

	t.Run("say 'Hello World', if was passed empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		verifyMessage(t, got, want)
	})
}
