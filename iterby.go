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

// Count generates an infinite sequence of ints.
func Count() func(func(int) bool) {
	return func(yield func(int) bool) {
		for i := 0; ; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

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

func RangeFilter[T any](begin func(T) bool, end func(T) bool, f func(func(T) bool)) func(func(T) bool) {
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
