package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	result := add(10, 5)
	expected := 15

	if result != expected {
		t.Errorf("add(10, 5) = %d; want %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := subtract(10, 5)
	expected := 5

	if result != expected {
		t.Errorf("subtract(10, 5) = %d; want%d", result, expected)

	}
}

func TestDoMath(t *testing.T) {
	resultAdd := doMath(10, 5, add)
	expectedAdd := 15
	if resultAdd != expectedAdd {
		log.Fatalf("Error - got %d but want %d", resultAdd, expectedAdd)
	}

}
