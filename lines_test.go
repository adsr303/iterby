package iterby_test

import (
	"fmt"
	"strings"

	"github.com/adsr303/iterby"
)

func ExampleFilterLines() {
	rangedLines := `
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
