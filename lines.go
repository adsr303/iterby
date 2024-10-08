package iterby

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"regexp"
)

// IterateLines returns an iterator over lines from r.
// It calls errHandler when there was an error while scanning r.
// It returns a single-use iterator.
func IterateLines(r io.Reader, errHandler func(error)) iter.Seq[string] {
	return func(yield func(string) bool) {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				return
			}
		}
		if err := scanner.Err(); err != nil {
			errHandler(err)
		}
	}
}

// A LineFilter scans over the provided [io.Reader] pushing the lines that
// fall within inclusive ranges from lines matching First to ones matching
// Last. Use [NewLineFilter] for simpler construction.
type LineFilter struct {
	First, Last *regexp.Regexp
}

// NoOpHandler ignores the error passed as argument.
// It can be used in [IterateLines] or [LineFilter.Iterate].
func NoOpHandler(_ error) {}

// NewLineFilter returns a [LineFilter] over lines from r that fall within
// inclusive ranges between regular expressions specified in first and last.
//
// It returns an error if first or last is not a valid regular expression as
// defined in [regexp/syntax].
func NewLineFilter(first, last string) (LineFilter, error) {
	rf, err := regexp.Compile(first)
	if err != nil {
		return LineFilter{}, fmt.Errorf("invalid regexp for first: %w", err)
	}
	rl, err := regexp.Compile(last)
	if err != nil {
		return LineFilter{}, fmt.Errorf("invalid regexp for last: %w", err)
	}
	return LineFilter{First: rf, Last: rl}, nil
}

// Iterate returns an iterator over lines filtered from r that fall within
// inclusive ranges between first and last provided to [NewLineFilter].
// It calls errHandler when there was an error while scanning r.
//
// It returns a single-use iterator.
func (f LineFilter) Iterate(r io.Reader, errHandler func(error)) iter.Seq[string] {
	i := IterateLines(r, errHandler)
	return RangeFilter(f.First.MatchString, f.Last.MatchString, i)
}
