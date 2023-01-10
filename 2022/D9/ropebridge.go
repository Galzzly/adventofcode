package main

import (
	"flag"
	"fmt"
	"image"
	"time"

	"github.com/Galzzly/adventofcode/utilities"
)

/*
	Read the input file, and parse it using the utility function to split the data into a useable slice.
	For Part 1, use the tvisit map to track which points have been visited by setting the key to true.
	This would mean that the point will only be added once, and additional visits will not increment the
	count more than necessary.
	For Part 2, use the tail map to track which points the end of the tope has visited. Uses the same
	method as part 1, using a bool value to indicate a point has been visited.
*/

var movement = map[rune]image.Point{
	'U': {0, -1},
	'R': {1, 0},
	'D': {0, 1},
	'L': {-1, 0},
}

func main() {
	start := time.Now()
	var file string
	flag.StringVar(&file, "input", "input.txt", "")
	flag.Parse()

	motions := utilities.ReadFileLineByLine(file)

	p1, p2 := solve(motions)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
	fmt.Println("Total time:", time.Since(start))
}

func solve(motions []string) (r1, r2 int) {
	tvisit := map[image.Point]bool{}
	tail := map[image.Point]bool{}
	rope := make([]image.Point, 10)
	for _, motion := range motions {
		var d rune
		var n int
		fmt.Sscanf(motion, "%c %d", &d, &n)
		for i := 0; i < n; i++ {
			rope[0] = rope[0].Add(movement[d])
			for i := 1; i < len(rope); i++ {
				if dir := rope[i-1].Sub(rope[i]); utilities.Abs(dir.X) > 1 || utilities.Abs(dir.Y) > 1 {
					rope[i] = rope[i].Add(image.Point{seeker(dir.X), seeker(dir.Y)})
				}
			}
			tvisit[rope[1]] = true
			tail[rope[len(rope)-1]] = true
		}
	}
	r1 = len(tvisit)
	r2 = len(tail)
	return
}

// seeker returns 1, -1, or 0 in order for the end of the rope to step a single value.
func seeker(n int) int {
	if n > 0 {
		return 1
	} else if n < 0 {
		return -1
	}
	return 0
}
