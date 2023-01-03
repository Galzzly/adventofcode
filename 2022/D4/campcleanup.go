package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/Galzzly/adventofcode/utilities"
)

/*
	Read the file in, and parse the input into the lines.
	Iterate over the lines, and calculate:
	Part 1, how many lines have pairs where one contains the other
	Part 2, how many lines have pairs where they overlap

	Make use of Sscanf to read the values from the input line, to save
	time converting string to int afterwards.
*/

func main() {
	start := time.Now()
	var file string
	flag.StringVar(&file, "input", "input.txt", "")
	flag.Parse()

	var p1, p2 int
	for _, line := range utilities.ReadFileLineByLine(file) {
		var e1start, e1end, e2start, e2end int
		fmt.Sscanf(line, "%d-%d,%d-%d", &e1start, &e1end, &e2start, &e2end)
		if e2start >= e1start && e2end <= e1end || e1start >= e2start && e1end <= e2end {
			p1++
		}
		if e2start <= e1end && e2end >= e1start || e1start <= e2end && e1end >= e2end {
			p2++
		}
	}
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
	fmt.Println("Total Time:", time.Since(start))
}
