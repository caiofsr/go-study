package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	iterations := Iterate("a", 5)
	want := "aaaaa"

	if iterations != want {
		t.Errorf("want '%s', got '%s'", want, iterations)
	}
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterate("a", 5)
	}
}

func ExampleIterate() {
	iteration := Iterate("a", 5)
	fmt.Println(iteration)
	// Output: aaaaa
}
