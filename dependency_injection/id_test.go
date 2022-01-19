package main

import (
	"bytes"
	"testing"
)

func TestGreets(t *testing.T) {
	buffer := bytes.Buffer{}
	Greets(&buffer, "Caio")

	result := buffer.String()
	expected := "Hello, Caio"

	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}
