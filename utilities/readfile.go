package utilities

import (
	"os"
)

// ReadFile reads in the file provided, and splits into a string slice
// using the function provided.
func ReadFile[T any](file string, split func(f string) T) T {
	f, err := os.ReadFile(file)
	Check(err)
	return split(string(f))
}
