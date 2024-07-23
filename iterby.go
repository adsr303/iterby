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
