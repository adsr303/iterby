package iterby_test

import (
	"fmt"
	"strings"

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

const rangedLines = `
outside
{
inside
}
out 2
out 3

{
abc

xyz
}
`

func ExampleFilterLines() {
	for s := range iterby.FilterLines(`\{`, `\}`, strings.NewReader(rangedLines)) {
		fmt.Println(s)
	}
	// Output:
	// {
	// inside
	// }
	// {
	// abc
	//
	// xyz
	// }
}
