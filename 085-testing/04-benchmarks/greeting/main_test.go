package greeting

import (
	"fmt"
	"testing"
)

func TestG(t *testing.T) {
	got := Greet("Pany")
	want := "Hello, Pany!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Pany"))
	// Output:
	// Hello, Pany!
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Pany")
	}
}