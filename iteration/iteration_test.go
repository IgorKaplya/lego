package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	expected := "aaaaa"
	got := Repeat("a", 5)

	if expected != got {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("b", 10)
	}
}

func ExampleRepeat() {
	got := Repeat("a", 7)
	fmt.Println(got)
	// Output: aaaaaaa
}
