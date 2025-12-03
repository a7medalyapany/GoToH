package word

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	s := "Hello world hello"
	expected := 3
	result := Count(s)
	if result != expected {
		t.Errorf("Count(%q) = %d; want %d", s, result, expected)
	}
}
func TestUseCount(t *testing.T) {
	type testCases struct {
		input    string
		expected map[string]int
	}

	xs := []testCases{
		{
			input: "Hello world hello",
			expected: map[string]int{
				"Hello": 1,
				"hello": 1,
				"world": 1,
			},
		},
		{
			input: "foo bar foo baz",
			expected: map[string]int{
				"foo": 2,
				"bar": 1,
				"baz": 1,
			},
		},
	}

	for _, tc := range xs {
		result := UseCount(tc.input)
		for k, v := range tc.expected {
			if result[k] != v {
				t.Errorf("UseCount(%q)[%q] = %d; want %d", tc.input, k, result[k], v)
			}
		}
	}
}

func ExampleCount() {
	fmt.Println(Count("My name is alyapany"))
	// Output:
	// 4
}
func ExampleUseCount() {
	fmt.Println(UseCount("go go gophers"))
	// Output:
	// map[go:2 gophers:1]
}

func BenchmarkCount(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		Count("Benchmarking the Count function with some sample text to measure performance.")
	}
}
func BenchmarkUseCount(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		UseCount("Benchmarking the UseCount function with some sample text to measure performance of UseCount.")
	}
}