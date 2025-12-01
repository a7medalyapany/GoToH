package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(1, 2, 3, 4, 5)
	expected := 15

	if result != expected {
		t.Errorf("sum(1, 2, 3, 4, 5) = %d; want %d", result, expected)
	}
}