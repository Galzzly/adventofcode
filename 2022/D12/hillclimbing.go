package main

import (
	"flag"
	"fmt"
	"image"
	"math"
	"sync"
	"time"

	"github.com/Galzzly/adventofcode/utilities"
)

/*
	Take the input and parse the lines into a string slice.
	Use the slice to build up the grid, returning the start and end points
	as well as a slice of the lowest elevation points (a).
	Part 1 is looking for the fewest steps required to go from start to end
	Part 2 is looking for the fewest steps to go from one of the lowest
	elevation points to the end.
*/

// Set Grid to have the type of rune
type Grid struct {
	*utilities.Grid[rune]
}

// Use a custom Neighbours function for Grid to use
func (g Grid) Neighbours(p image.Point) []image.Point {
	val := g.GetState(p)
	return utilities.Select(g.Grid.Neighbours(p), func(x image.Point) bool {
		return g.GetState(x) <= val+1
	})
}

func main() {
	start := time.Now()
	var file string
	flag.StringVar(&file, "input", "input.txt", "")
	flag.Parse()

	lines := utilities.ReadFileLineByLine(file)
	grid, s, e, a := buildGrid(lines)

	t1 := time.Now()
	fmt.Println("Part 1:", len(utilities.Search(grid, s, e))-1, "- Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", solve(grid, a, e), "- Took:", time.Since(t2))
	fmt.Println("Total time:", time.Since(start))
}

/*
The solution for Part 2 makes use of a WaitGroup, channels, and goroutines
to test each lowest elevation points concurrently. Watching the output of
the channel and testing if it is lower that the current result.
*/
func solve(g Grid, a []image.Point, e image.Point) (res int) {
	res = math.MaxInt
	distance := make(chan int, len(a))
	var wg sync.WaitGroup
	for _, p := range a {
		wg.Add(1)
		go func(p image.Point) {
			defer wg.Done()
			d := len(utilities.Search(g, p, e)) - 1
			distance <- d
		}(p)
	}
	wg.Wait()
	close(distance)
	for d := range distance {
		if d < res && d > 1 {
			res = d
		}
	}
	return
}

func buildGrid(lines []string) (grid Grid, s, e image.Point, a []image.Point) {
	grid.Grid = utilities.NewGrid[rune](len(lines[0]), len(lines), utilities.CrossDelta)
	for y, line := range lines {
		for x, c := range line {
			switch c {
			case 'S':
				s = image.Point{x, y}
				grid.SetState(x, y, 'a')
			case 'E':
				e = image.Point{x, y}
				grid.SetState(x, y, 'z')
			case 'a':
				a = append(a, image.Point{x, y})
				grid.SetState(x, y, 'a')
			default:
				grid.SetState(x, y, c)
			}

		}
	}
	return
}
