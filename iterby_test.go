//go:build goexperiment.rangefunc

package iterby_test

import (
	"fmt"

	"github.com/adsr303/iterby"
)

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
