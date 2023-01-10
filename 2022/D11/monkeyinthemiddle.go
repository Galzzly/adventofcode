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
	Read in the input file, and parse it to getMonkeys to get the initial outline of the items
	that they are each holding, and to set up the operation & test for each monkey.
	The LCM is calculated for the Part 2 solution.
*/

type Monkey struct {
	items []int
	op    func(int) int
	test  func(int) int
}

func main() {
	start := time.Now()
	var file string
	flag.StringVar(&file, "input", "input.txt", "")
	flag.Parse()

	monkeys, lcm := getMonkeys(utilities.ReadFileDoubleLine(file))

	t1 := time.Now()
	fmt.Println("Part 1:", solve(monkeys, func(i int) int { return i / 3 }, 20), "- Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", solve(monkeys, func(i int) int { return i % lcm }, 10000), "- Took:", time.Since(t2))

	fmt.Println("Total time:", time.Since(start))
}

func solve(monkeys []Monkey, translation func(i int) int, rounds int) (res int) {
	monkeys = append([]Monkey{}, monkeys...)
	inspects := make([]int, len(monkeys))
	for i := 0; i < rounds; i++ {
		for m := 0; m < len(monkeys); m++ {
			for _, item := range monkeys[m].items {
				item = translation(monkeys[m].op(item))
				nm := monkeys[m].test(item)
				monkeys[nm].items = append(monkeys[nm].items, item)
				inspects[m]++
			}
			monkeys[m].items = nil
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspects)))
	res = inspects[0] * inspects[1]
	return
}

func getMonkeys(lines []string) (monkeys []Monkey, lcm int) {
	monkeys = make([]Monkey, len(lines))
	lcm = 1
	for m, l := range lines {
		line := strings.Split(l, "\n")
		// fmt.Println(line[2])
		monkeys[m].items = getItems(line[1])
		operation := strings.Fields((strings.ReplaceAll(strings.Split(line[2], " = ")[1], "* old", "^ 2")))
		// fmt.Println(operation[2])
		switch operation[1] {
		case "+":
			monkeys[m].op = func(i int) int { return i + utilities.Atoi(operation[2]) }
		case "*":
			monkeys[m].op = func(i int) int { return i * utilities.Atoi(operation[2]) }
		case "^":
			monkeys[m].op = func(i int) int { return i * i }
		}
		test := utilities.Atoi(strings.Fields(line[3])[3])
		monkeys[m].test = func(i int) int {
			if i%test == 0 {
				return utilities.Atoi(strings.Fields(line[4])[5])
			}
			return utilities.Atoi(strings.Fields(line[5])[5])
		}
		lcm *= test
	}
	return
}

func getItems(line string) (items []int) {
	l := strings.Split(line, ": ")[1]
	for _, item := range strings.Split(l, ", ") {
		items = append(items, utilities.Atoi(item))
	}
	return
}
