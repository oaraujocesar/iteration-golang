package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeat := Repeat("a", 6)
	expected := "aaaaaa"

	if repeat != expected {
		t.Errorf("expected %q but got %q", expected, repeat)
	}
}

func ExampleRepeat() {
	repeat := Repeat("a", 7)
	fmt.Println(repeat)
	// Output: aaaaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 15)
	}
}
