package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}
	want := 6

	got := Sum(numbers)

	if got != want {
		t.Errorf("got %d want %d given %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	want := []int{3, 9}

	got := SumAll([]int{1, 2}, []int{0, 9})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got []int, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sum of some slices", func(t *testing.T) {
		want := []int{2, 9}
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		want := []int{0, 9}
		got := SumAllTails([]int{}, []int{0, 9})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
