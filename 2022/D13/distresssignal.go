package main

import (
	"flag"
	"fmt"
	"sort"
	"time"

	"github.com/Galzzly/adventofcode/utilities"
)

/*
Read the input file, and parse it to create the inputs needed for parts 1 and 2
(i1 and i2).
Part 1 requires the sum of the index of pairs that are already in the correct order.
Part 2 requires all of the packets to be in the correct order, and for the index of
[[2]] and [[6]] to be multiplied together.
*/

func main() {
	start := time.Now()
	var file string
	flag.StringVar(&file, "input", "input.txt", "")
	flag.Parse()

	i1, i2 := getInput(utilities.ReadFileDoubleSingle(file))

	t1 := time.Now()
	fmt.Println("Part 1:", part1(i1), "- Took", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(i2), "- Took", time.Since(t2))

	fmt.Println("Total time:", time.Since(start))
}

func part1(input [][]any) (res int) {
	for k, pair := range input {
		if r := compare(pair[0], pair[1]); r >= 0 {
			res += k + 1
		}
	}
	return
}

func part2(input []any) (res int) {
	sort.Slice(input, func(i, j int) bool {
		a := input[i]
		b := input[j]
		return compare(a, b) == 1
	})
	res = 1
	for i, p := range input {
		v := fmt.Sprintf("%v", p)
		if v == "[[[2]]]" || v == "[[[6]]]" {
			res *= i + 1
		}
	}
	return
}

/*
compare takes in the two sides of the pair, and compares the value
First checks if the items are both integers, and returns 1, -1, or 0 based
on the test case.
Then checks if the itmes are lists, if an item is not a list, it will be created
as one with the appropriate integer value added.
The list lengths are compared, if the end of the left side is reached then 1 is
returned, if the end of the right side is reached then -1 is returned, otherwise
compare both sides again.
*/
func compare(l, r any) int {
	lInt, lIsInt := l.(int)
	rInt, rIsInt := r.(int)
	if lIsInt && rIsInt {
		switch {
		case lInt < rInt:
			return 1
		case lInt > rInt:
			return -1
		default:
			return 0
		}
	}

	lList, lIsList := l.([]any)
	rList, rIsList := r.([]any)
	if !lIsList {
		lList = []any{lInt}
	}
	if !rIsList {
		rList = []any{rInt}
	}

	lLen, rLen := len(lList), len(rList)
	max := utilities.Biggest(lLen, rLen)
	for i := 0; i < max; i++ {
		if i >= lLen {
			return 1
		}
		if i >= rLen {
			return -1
		}
		if sub := compare(lList[i], rList[i]); sub != 0 {
			return sub
		}
	}
	return 0
}
func getInput(lines [][]string) (i1 [][]any, i2 []any) {
	for _, line := range lines {
		pair := getPair(line)
		i2 = append(i2, pair...)
		i1 = append(i1, pair)
	}
	pair := getPair([]string{"[[2]]", "[[6]]"})
	i2 = append(i2, pair...)
	return
}

func getPair(lines []string) (pair []any) {
	pair = make([]any, 0, len(lines))
	for _, line := range lines {
		np, _ := parseLine(line)
		pair = append(pair, np)
	}
	return
}

func parseLine(line string) (any, int) {
	out := make([]any, 0)
	nC := make([]rune, 0)
	var i int
	for i = 0; i < len(line); i++ {
		c := line[i]
		switch c {
		case '[':
			o, k := parseLine(line[i+1:])
			out = append(out, o)
			i += k
		case ']':
			if len(nC) > 0 {
				n := utilities.Atoi(string(nC))
				out = append(out, n)
			}
			return out, i + 1
		case ',':
			if len(nC) > 0 {
				n := utilities.Atoi(string(nC))
				out = append(out, n)
				nC = make([]rune, 0)
			}
		default:
			nC = append(nC, rune(c))
		}
	}
	return out, i
}
