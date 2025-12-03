package dog

import (
	"fmt"
	"testing"
)


func TestYears(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{1, 7},
		{2, 14},
		{3, 21},
		{10, 70},
		{0, 0},
		{-5, -35},
	}

	for _, test := range tests {
		result := Years(test.input)
		if result != test.expected {
			t.Errorf("Years(%d) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func ExampleYears() {
	fmt.Println(Years(3))
	// Output:
	// 21
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}