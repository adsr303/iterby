package iterby_test

import (
	"fmt"
	"strings"

	"github.com/adsr303/iterby"
)

func ExampleLineFilter() {
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
	f, err := iterby.NewLineFilter(`\{`, `\}`)
	if err != nil {
		fmt.Println(err)
	}
	for s := range f.Iterate(strings.NewReader(rangedLines), iterby.NoOpHandler) {
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
