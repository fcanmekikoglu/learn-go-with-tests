package integers

import (
	"fmt"
	"testing"
)

func TestExtract(t *testing.T) {
	extraction := Extract(4, 2)
	expected := 2
	if extraction != expected {
		t.Errorf("expected '%d' but got '%d'", expected, extraction)
	}
}

func ExampleExtract() {
	extraction := Extract(5, 1)
	fmt.Println(extraction)
	// Output: 4
}
