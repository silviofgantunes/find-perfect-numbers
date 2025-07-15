package service

import (
	"reflect"
	"sort"
	"testing"
)

func TestIsPerfect(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{6, true},
		{28, true},
		{496, true},
		{8128, true},
		{1, false},
		{0, false},
		{-6, false},
		{12, false},
		{33550336, true}, // The 5th perfect number
	}

	for _, tt := range tests {
		t.Run("IsPerfect", func(t *testing.T) {
			result := IsPerfect(tt.input)
			if result != tt.expected {
				t.Errorf("IsPerfect(%d) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFindPerfectNumbers(t *testing.T) {
	result := FindPerfectNumbers(1, 10000)
	expected := []int{6, 28, 496, 8128}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FindPerfectNumbers(1, 10000) = %v; expected %v", result, expected)
	}
}

func TestFindPerfectNumbersParallel(t *testing.T) {
	result := FindPerfectNumbersParallel(1, 10000)
	expected := []int{6, 28, 496, 8128}

	// Parallel execution may result in unordered output
	sort.Ints(result)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FindPerfectNumbersParallel(1, 10000) = %v; expected %v", result, expected)
	}
}

func TestSequentialAndParallelMatch(t *testing.T) {
	seq := FindPerfectNumbers(1, 100000)
	par := FindPerfectNumbersParallel(1, 100000)

	sort.Ints(seq)
	sort.Ints(par)

	if !reflect.DeepEqual(seq, par) {
		t.Errorf("Mismatch between sequential and parallel results:\nSequential: %v\nParallel: %v", seq, par)
	}
}
