package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/Galzzly/adventofcode/utilities"
)

/*
	Read the input file, and parse it using the utility function to split the data into a usable
	slice.
	Play the game, and make use of the results map, that has been generated using the instructions
	from https://adventofcode.com/2022/day/2
*/

var results = map[string][]int{
	"A X": {4, 3},
	"A Y": {8, 4},
	"A Z": {3, 8},
	"B X": {1, 1},
	"B Y": {5, 5},
	"B Z": {9, 9},
	"C X": {7, 2},
	"C Y": {2, 6},
	"C Z": {6, 7},
}

func main() {
	start := time.Now()
	var file string
	flag.StringVar(&file, "input", "input.txt", "")
	flag.Parse()

	input := utilities.ReadFile(file, func(f string) []string { return strings.Split(strings.TrimSpace(f), "\n") })
	r1, r2 := game(input)
	fmt.Println("Part 1:", r1)
	fmt.Println("Part 2:", r2)
	fmt.Println("Total time:", time.Since(start))
}

func game(input []string) (r1, r2 int) {
	for _, v := range input {
		r1 += results[v][0]
		r2 += results[v][1]
	}
	return
}
