package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/Galzzly/adventofcode/utilities"
)

/*
	Read in the input file, and parse it using the utility function to split the data into
	useable slices.
	Use the countCalories function to calculate the colories that each elf is carrying, and
	reverse sort the results to allow for quicker processing afterwards.
	Part 1 requires the total amount of calories being carried by the elf carrying the most.
	Part 2 requires the total amount of calories being carried by the top three elves.
*/

func main() {
	start := time.Now()
	var file string
	flag.StringVar(&file, "input", "input.txt", "")
	flag.Parse()

	elves := utilities.ReadFile(file, func(f string) [][]string {
		res := [][]string{}
		for _, line := range strings.Split(strings.TrimSpace(f), "\n\n") {
			res = append(res, strings.Split(line, "\n"))
		}
		return res
	})

	calories := countCalories(elves)
	t1 := time.Now()
	fmt.Println("Part 1:", calories[0], "- Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", utilities.SumArray(calories[:3]), "- Took:", time.Since(t2))
	fmt.Println("Total time:", time.Since(start))
}

func countCalories(elves [][]string) (res []int) {
	res = make([]int, len(elves))
	for i := range elves {
		for _, c := range elves[i] {
			res[i] += utilities.Atoi(c)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(res)))
	return
}
