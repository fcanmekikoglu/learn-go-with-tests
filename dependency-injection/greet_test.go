package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}

	})
}
