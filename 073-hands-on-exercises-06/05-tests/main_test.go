package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(5, 5)

	if total != 10 {
		t.Errorf("add(5, 5) = %d; want 10", total)
	}
}
