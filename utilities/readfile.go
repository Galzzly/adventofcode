package utilities

import (
	"os"
	"strings"
)

/*
ReadFile reads in the file provided, and splits into a string slice
using the function provided.
This will be the parent function, that can still be called for custom
splits
*/
func ReadFile[T any](file string, split func(f string) T) T {
	f, err := os.ReadFile(file)
	Check(err)
	return split(string(f))
}

func ReadFileLineByLine(file string) []string {
	return ReadFile(file, func(f string) []string {
		return strings.Split(strings.TrimSpace(f), "\n")
	})
}

func ReadFileDoubleLine(file string) []string {
	return ReadFile(file, func(f string) []string {
		return strings.Split(strings.TrimSpace(f), "\n\n")
	})
}

func ReadFileDoubleSingle(file string) [][]string {
	return ReadFile(file, func(f string) [][]string {
		res := [][]string{}
		for _, line := range strings.Split(strings.TrimSpace(f), "\n\n") {
			res = append(res, strings.Split(line, "\n"))
		}
		return res
	})
}
