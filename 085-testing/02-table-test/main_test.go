package main

import "testing"

func TestSum(t *testing.T) {
	type testCases struct {
		data    []int
		answer int
	}

	tests := []testCases{
		testCases{data: []int{1, 2, 3, 4, 5}, answer: 15},
		testCases{data: []int{10, 20, 30}, answer: 60},
		testCases{data: []int{-1, -2, -3, -4, 4}, answer: -6},
	}

	for _, tc := range tests {
		result := sum(tc.data...)
		if result != tc.answer {
			t.Errorf("sum(%v) = %d; want %d", tc.data, result, tc.answer)
		}
	}
}