package iterby_test

import (
	"fmt"
	"iter"
	"slices"
	"testing"

	"github.com/adsr303/iterby"
)

func TestEnumerate(t *testing.T) {
	s := []string{"a", "b", "c"}
	seq := func() iter.Seq[string] {
		return func(yield func(string) bool) {
			for _, x := range s {
				if !yield(x) {
					return
				}
			}
		}
	}
	for i, x := range iterby.Enumerate(seq()) {
		if x != s[i] {
			t.Errorf("expected s[%d] to be %s, got %s", i, s[i], x)
		}
	}
}

func TestChain(t *testing.T) {
	s := make([]string, 0)
	for x := range iterby.Chain([]string{"foo", "bar"}, []string{"baz"}) {
		s = append(s, x)
	}
	expected := []string{"foo", "bar", "baz"}
	if !slices.Equal(s, expected) {
		t.Errorf("expected %v, got %v", expected, s)
	}
}

func ExampleCount() {
	for i := range iterby.Count() {
		if i > 5 {
			break // Stop the test at some point :)
		}
		fmt.Println(i)
	}
	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
}

func ExampleCount2() {
	for x := range iterby.Count2(0, 3.14/8) {
		if x > 7 {
			break // Stop the test at some point :)
		}
		fmt.Printf("%.2f ", x)
	}
	// Output:
	// 0.00 0.39 0.79 1.18 1.57 1.96 2.35 2.75 3.14 3.53 3.93 4.32 4.71 5.10 5.50 5.89 6.28 6.67
}
