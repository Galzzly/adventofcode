package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/Galzzly/adventofcode/utilities"
)

/*
	Read in the input file, and parse it using the utility function to grab the useful string.
	The solution for both parts is to parse the string and look for a number of unique characters
	and returning the character ID after the last of the unique characters.
*/

func main() {
	start := time.Now()
	var file string
	flag.StringVar(&file, "input", "input.txt", "")
	flag.Parse()

	stream := utilities.ReadFile(file, func(f string) string {
		return strings.Split(strings.TrimSpace(f), "\n")[0]
	})

	t1 := time.Now()
	fmt.Println("Part 1:", streamcheck(stream, 4), "- Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", streamcheck(stream, 14), "- Took:", time.Since(t2))
	fmt.Println("Total time: ", time.Since(start))
}

func streamcheck(stream string, length int) int {
	for i := 0; i < len(stream)-length; i++ {
		if checkUnique(stream[i : i+length]) {
			return i + length
		}
	}
	return 0
}

func checkUnique(s string) bool {
	for _, c := range s {
		if strings.Count(s, string(c)) > 1 {
			return false
		}
	}
	return true
}
