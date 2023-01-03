package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/Galzzly/adventofcode/utilities"
)

/*
	Read in the input file, and parse it using the utility function to split the data into
	useable slices.
	Part 1 iterates through the backpacks, which have been split into equal parts using
	the getRucksacks function, checking if an item is present in both slices.
	Part 2 iterates through backpacks in groups of three, to find items that are common
	between all three.
*/

var values = map[rune]int{
	'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7,
	'h': 8, 'i': 9, 'j': 10, 'k': 11, 'l': 12, 'm': 13, 'n': 14,
	'o': 15, 'p': 16, 'q': 17, 'r': 18, 's': 19, 't': 20, 'u': 21,
	'v': 22, 'w': 23, 'x': 24, 'y': 25, 'z': 26,
	'A': 27, 'B': 28, 'C': 29, 'D': 30, 'E': 31, 'F': 32, 'G': 33,
	'H': 34, 'I': 35, 'J': 36, 'K': 37, 'L': 38, 'M': 39, 'N': 40,
	'O': 41, 'P': 42, 'Q': 43, 'R': 44, 'S': 45, 'T': 46, 'U': 47,
	'V': 48, 'W': 49, 'X': 50, 'Y': 51, 'Z': 52,
}

func main() {
	start := time.Now()
	var file string
	flag.StringVar(&file, "input", "input.txt", "")
	flag.Parse()

	lines := utilities.ReadFile(file, func(f string) []string { return strings.Split(strings.TrimSpace(f), "\n") })

	t1 := time.Now()
	fmt.Println("Part 1:", part1(lines), "- Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(lines), "- Took:", time.Since(t2))
	fmt.Println("Total time:", time.Since(start))
}

func part1(lines []string) (r1 int) {
bag:
	for _, r := range getRucksacks(lines) {
		for _, c := range r[0] {
			if strings.ContainsRune(r[1], c) {
				r1 += values[c]
				continue bag
			}
		}
	}
	return
}

func getRucksacks(lines []string) (r [][]string) {
	for _, line := range lines {
		l := len(line) / 2
		r = append(r, []string{line[:l], line[l:]})
	}
	return
}

func part2(lines []string) (r2 int) {
bags:
	for i := 0; i < len(lines); i += 3 {
		for _, c := range lines[i] {
			if strings.ContainsRune(lines[i+1], c) && strings.ContainsRune(lines[i+2], c) {
				r2 += values[c]
				continue bags
			}
		}
	}
	return
}
