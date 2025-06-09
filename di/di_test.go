package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	want := "Hello, Chris"
	got := bytes.Buffer{}

	Greet(&got, "Chris")

	if got.String() != want {
		t.Errorf("got %q want %q", got, want)
	}
}
