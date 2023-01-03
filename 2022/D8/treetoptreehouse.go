package main

import (
	"flag"
	"fmt"
	"image"
	"sync"
	"time"

	"github.com/Galzzly/adventofcode/utilities"
)

/*
	Read the input file, and transform it into a map using the image.Point as the
	key.
	Part 1 looks for the number of trees visible from the outside of the grid
	Part 2 looks for the highest scenic score for a tree

	Concurrency is used to reduce the time spent waiting, and channels are used to
	take the output from lookarounds.

	The lookaround makes use of a Delta variable defined in utilities, which will
	change the point per the direction until the outside of the grid is reached.
*/

type Treemap map[image.Point]int

func main() {
	start := time.Now()
	var file string
	flag.StringVar(&file, "input", "input.txt", "")
	flag.Parse()

	treemap := utilities.MakeIntImagePoint(utilities.ReadFileLineByLine(file))

	p1, p2 := solve(treemap)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
	fmt.Println("Total time:", time.Since(start))
}

func solve(treemap Treemap) (r1, r2 int) {
	vis := make(chan int, len(treemap))
	score := make(chan int, len(treemap))
	var wg sync.WaitGroup
	wg.Add(len(treemap))
	go func() {
		for k, v := range treemap {
			go treemap.lookaround(k, v, vis, score, &wg)
		}
		wg.Wait()
		close(vis)
		close(score)
	}()
	for v := range vis {
		r1 += v
	}
	for s := range score {
		if s > r2 {
			r2 = s
		}
	}
	return
}

func (t Treemap) lookaround(k image.Point, v int, vis, score chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	r1, r2 := 0, 1
	for _, p := range utilities.CrossDelta {
		for np, i := k.Add(p), 0; ; np, i = np.Add(p), i+1 {
			if _, ok := t[np]; !ok {
				r1, r2 = 1, r2*i
				break
			}
			if t[np] >= v {
				r2 *= i + 1
				break
			}
		}
	}
	vis <- r1
	score <- r2
}
