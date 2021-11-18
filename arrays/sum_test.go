package main

import (
	"reflect"
	"testing"
)

// To get verbose output, run the test with -v flag
// To get coverage results, run the test with -cover flag

func TestSum(t *testing.T) {

	// When you dont specify the size of the array, you get a slice
	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	want := 6

	if got != want {
		t.Errorf("got %d want %d, given %v", got, want, numbers)
	}

}

func TestSumArray(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := SumArray(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d, given %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	sliceOne := []int{1, 2, 3}
	sliceTwo := []int{4, 5, 6}

	got := SumAll(sliceOne, sliceTwo)
	want := []int{6, 15}

	// Note that reflect.DeepEqual is not type safe
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 10})
		want := []int{5, 10}

		checkSums(t, got, want)
	})

	t.Run("sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{})
		want := []int{2, 0}

		checkSums(t, got, want)
	})

}

func TestSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	checkResult := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("slice low only", func(t *testing.T) {

		got := slice[1:]
		want := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}

		checkResult(t, got, want)
	})

	t.Run("slice high only", func(t *testing.T) {
		got := slice[:5]
		want := []int{1, 2, 3, 4, 5}
		checkResult(t, got, want)
	})

	t.Run("slice both ends", func(t *testing.T) {
		got := slice[4:6]
		want := []int{5, 6}
		checkResult(t, got, want)
	})

	t.Run("slice both both end to understand", func(t *testing.T) {
		got := slice[1:2]
		// Just like Java's String.substring(low, high)
		want := []int{2}
		checkResult(t, got, want)
	})
}
