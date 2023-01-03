package main

import (
	"flag"
	"fmt"
	"path"
	"strings"
	"time"
	"unicode"

	"github.com/Galzzly/adventofcode/utilities"
)

/*
	Read the input file, and parse it using the utility function to split the data into a useable slice.
	Generate the filesystem using the getFs function (see explanation below), and solve by iterating
	over the generated filesystem, adding the size to the total if it is above m1 for part 1, and finding
	the minimum directory size that will provide m2 for part 2.
*/

type FS map[string]int

func main() {
	start := time.Now()
	var file string
	flag.StringVar(&file, "input", "input.txt", "")
	flag.Parse()

	lines := utilities.ReadFileLineByLine(file)

	p1, p2 := getFs(lines).solve(100000, 30000000)

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
	fmt.Println("Total time:", time.Since(start))
}

func (f FS) solve(m1, m2 int) (r1, r2 int) {
	r2 = f["/"]
	for _, s := range f {
		if s <= m1 {
			r1 += s
		}
		if s+70000000-f["/"] >= m2 && s < r2 {
			r2 = s
		}
	}
	return
}

/*
getFs iterates over the lines that are read in, and returns the filesystem properties
as a map, using the path as the key, and containing the size.
When a number is found in the first field, the size is added to all of the parent
directories sizes.
*/
func getFs(lines []string) FS {
	fs := make(FS)
	var p string
	for _, line := range lines {
		if strings.HasPrefix(line, "$ cd") {
			p = path.Join(p, strings.Fields(line)[2])
		} else if unicode.IsDigit([]rune(line)[0]) {
			var size int
			fmt.Sscanf(line, "%d", &size)
			for d := p; d != "/"; d = path.Dir(d) {
				fs[d] += size
			}
			fs["/"] += size
		}
	}

	return fs
}
