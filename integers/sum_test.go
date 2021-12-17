package integers

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	sum := Add(2, 2)
	want := 4

	if sum != want {
		t.Errorf("want '%d', sum '%d'", want, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
