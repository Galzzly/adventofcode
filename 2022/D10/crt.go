package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/Galzzly/adventofcode/utilities"
)

/*
	Read the input file and parse it usint the utility function so that the instructions are
	contained in a slice.
	Making use of Sscanf to iterate over the instructions, use the instr to define what to do
	and what to append where.
*/

func main() {
	start := time.Now()
	var file string
	flag.StringVar(&file, "input", "input.txt", "")
	flag.Parse()

	lines := utilities.ReadFileLineByLine(file)
	p1, p2 := solve(lines)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:")
	for _, r := range p2 {
		fmt.Println(strings.Join(r, ""))
	}
	fmt.Println("Total time:", time.Since(start))
}

func solve(lines []string) (r1 int, r2 [][]string) {
	x := 1
	var cycle, c int
	row := []string{}
	for _, line := range lines {
		var instr string
		var num int
		fmt.Sscanf(line, "%s %d", &instr, &num)
		switch instr {
		case "noop":
			if c >= x-1 && c <= x+1 {
				row = append(row, "#")
			} else {
				row = append(row, " ")
			}
			cycle++
			c++
			switch cycle {
			case 20, 60, 100, 140, 180, 220:
				r1 += cycle * x
			case 40, 80, 120, 160, 200, 240:
				r2 = append(r2, row)
				row = []string{}
				c = 0
			}
		case "addx":
			for i := 0; i < 2; i++ {
				if c >= x-1 && c <= x+1 {
					row = append(row, "#")
				} else {
					row = append(row, " ")
				}
				cycle++
				c++
				switch cycle {
				case 20, 60, 100, 140, 180, 220:
					r1 += cycle * x
				case 40, 80, 120, 160, 200, 240:
					r2 = append(r2, row)
					row = []string{}
					c = 0
				}
			}
			x += num
		}
	}
	return
}
