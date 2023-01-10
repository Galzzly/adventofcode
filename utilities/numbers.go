package utilities

import "strconv"

// Atoi converts string to int, only to remove the error
// checking from the AOC code for the day
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}

// SumArray takes in an int slice, and returns its sum
func SumArray(a []int) int {
	var sum int
	for _, v := range a {
		sum += v
	}
	return sum
}

// Abs returns the absolute value of the input
func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Biggest(a, b int) (res int) {
	if a < b {
		return b
	}
	return a
}
