//go:build goexperiment.rangefunc

package iterby

// Enumerate allows to iterate over single-value rangefunc with indexes,
// similar to for-range over slices:
//
//	for i, x := range iterby.Enumerate(f) {
//	    fmt.Println(i, x)
//	}
func Enumerate[T any](f func(func(T) bool)) func(func(int, T) bool) {
	return func(yield func(int, T) bool) {
		var index int
		for t := range f {
			if !yield(index, t) {
				return
			}
			index++
		}
	}
}

// Count generates an infinite sequence of consecutive integers,
// starting from 0.
func Count() func(func(int) bool) {
	return func(yield func(int) bool) {
		for i := 0; ; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// Number is any integer or floating-point numeric type.
type Number interface {
	~int | ~uint | ~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32 | ~int64 | ~uint64 | ~float32 | ~float64
}

// Count2 generates an infinite sequence of numbers, starting from start
// and incrementing by step.
func Count2[T Number](start, step T) func(func(T) bool) {
	return func(yield func(T) bool) {
		for i := start; ; i += step {
			if !yield(i) {
				return
			}
		}
	}
}

// Chain generates a sequence of all elements of provided slices,
// allowing to for-range over them in a single loop.
func Chain[T any](args ...[]T) func(func(T) bool) {
	return func(yield func(T) bool) {
		for _, s := range args {
			for _, t := range s {
				if !yield(t) {
					return
				}
			}
		}
	}
}

// Chain generates an infinitely repeating sequence of all elements of
// provided slices, allowing to for-range over them in a single loop.
func Cycle[T any](args ...[]T) func(func(T) bool) {
	return func(yield func(T) bool) {
		for {
			for t := range Chain(args...) {
				if !yield(t) {
					return
				}
			}
		}
	}
}

// RangeFilter matches consecutive elements generated by f as predicated by
// begin and end. A range starts with the element for which begin returns true
// and ends with the element for which end returns true.
//
// RangeFilter was inspired by [AWK's record ranges] and [Perl's range operators].
//
// [AWK's record ranges]: https://www.gnu.org/software/gawk/manual/html_node/Ranges.html
// [Perl's range operators]: https://perldoc.perl.org/perlop#Range-Operators
func RangeFilter[T any](begin, end func(T) bool, f func(func(T) bool)) func(func(T) bool) {
	return func(yield func(T) bool) {
		var inRange bool
		for t := range f {
			if !inRange {
				if begin(t) {
					inRange = true
				} else {
					continue
				}
			}
			if !yield(t) {
				return
			}
			if end(t) {
				inRange = false
			}
		}
	}
}
