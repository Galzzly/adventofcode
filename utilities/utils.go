package utilities

// Check is a simple error checking function that panics if
// there is an error.
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Select returns a slice of items in the input that satisfy the test
// provided.
func Select[T any](in []T, f func(i T) bool) (res []T) {
	res = make([]T, 0)
	for _, v := range in {
		if f(v) {
			res = append(res, v)
		}
	}
	return
}
