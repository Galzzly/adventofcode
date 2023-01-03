package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/Galzzly/adventofcode/utilities"
)

/*
	Read in the input file, and parse it using the utility function to split the data into a
	usable slice, split by double line, and then by single line.
	Build the stacks used for both parts 1 and 2.
	Build up the rules to use by making use of Sscanf to extract the values in correct format.
	Solve the challenge by iterating through the rules.
*/

// Rule struct represents the number of items to move (num)
// from stack to stack (from, to)
type Rule struct {
	num      int
	from, to int
}

func main() {
	start := time.Now()
	var file string
	flag.StringVar(&file, "input", "input.txt", "")
	flag.Parse()

	sections := utilities.ReadFile(file, func(f string) (res [][]string) {
		res = [][]string{}
		for _, lines := range strings.Split(strings.TrimSpace(f), "\n\n") {
			res = append(res, strings.Split(lines, "\n"))
		}
		return
	})

	stack, keys := getStacks(sections[0])
	rules := getRules(sections[1])

	p1, p2 := solve(stack, rules, keys)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
	fmt.Println("Total Time: ", time.Since(start))
}

func solve(stack []string, rules []Rule, keys []int) (r1, r2 string) {
	// Copy the stack
	stack2 := make([]string, len(stack))
	copy(stack2, stack)
	for _, rule := range rules {
		for i := 0; i < rule.num; i++ {
			stack[rule.to] += string(stack[rule.from][len(stack[rule.from])-1])
			stack[rule.from] = stack[rule.from][:len(stack[rule.from])-1]
		}
		stack2[rule.to] += string(stack2[rule.from][len(stack2[rule.from])-rule.num:])
		stack2[rule.from] = stack2[rule.from][:len(stack2[rule.from])-rule.num]
	}
	for _, i := range keys {
		r1 += string(stack[i][len(stack[i])-1])
		r2 += string(stack2[i][len(stack2[i])-1])
	}
	return
}

/*
To build the stack, get the string of keys used by replacing all whitesapces to ""
and converting each to integer.
Iterate through the stack from the bottom up, and build up the the map.
*/
func getStacks(lines []string) (stack []string, k []int) {
	keys := lines[len(lines)-1]
	for _, n := range strings.ReplaceAll(keys, " ", "") {
		k = append(k, utilities.Atoi(string(n)))
	}
	stack = make([]string, len(k)+1)
	for i := len(lines) - 2; i >= 0; i-- {
		for j, c := range lines[i] {
			if unicode.IsLetter(c) {
				stack[utilities.Atoi(string(keys[j]))] += string(c)
			}
		}
	}
	return
}

func getRules(lines []string) (rules []Rule) {
	for _, rule := range lines {
		var a, b, c int
		fmt.Sscanf(rule, "move %d from %d to %d", &a, &b, &c)
		rules = append(rules, Rule{a, b, c})
	}
	return
}
