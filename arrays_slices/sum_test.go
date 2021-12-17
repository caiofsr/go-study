package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	verifySums := func(t *testing.T, got, wanted []int) {
		t.Helper()

		if !reflect.DeepEqual(got, wanted) {
			t.Errorf("result %v, wanted %v", got, wanted)
		}
	}

	t.Run("collection with any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		result := Sum(numbers)
		expected := 6

		if expected != result {
			t.Errorf("result was incorrect, got: %d, want: %d. given: %v", result, expected, numbers)
		}
	})

	t.Run("sum two slices and return one slice", func(t *testing.T) {
		result := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result was incorrect, got: %d, want: %d", result, expected)
		}
	})

	t.Run("faz as somas de alguns slices", func(t *testing.T) {
		result := SumAllRest([]int{1, 2}, []int{0, 9})
		wanted := []int{2, 9}

		verifySums(t, result, wanted)
	})

	t.Run("soma slices vazios de forma segura", func(t *testing.T) {
		result := SumAllRest([]int{}, []int{3, 4, 5})
		wanted := []int{0, 9}

		verifySums(t, result, wanted)
	})
}
