package iterby

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"os"
	"regexp"
)

// IterateLines returns an iterator over lines from r.
// It returns a single-use iterator.
func IterateLines(r io.Reader) iter.Seq[string] {
	return func(yield func(string) bool) {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				return
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprint(os.Stderr, err)
		}
	}
}

// FilterLines returns an iterator over lines from r that fall within ranges
// between (inclusive) regular expressions specified in begin and end.
//
// It panics if begin or end is not a valid regular expression as defined in
// [regexp/syntax].
//
// It returns a single-use iterator.
func FilterLines(begin, end string, r io.Reader) iter.Seq[string] {
	b := regexp.MustCompile(begin)
	e := regexp.MustCompile(end)
	i := IterateLines(r)
	return RangeFilter(b.MatchString, e.MatchString, i)
}
