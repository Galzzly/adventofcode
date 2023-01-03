package utilities

import (
	"image"
	"strings"
)

var CrossDelta = []image.Point{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

var FullDelta = []image.Point{
	{-1, -1}, {0, -1}, {1, -1}, // up-left, up, up-right
	{-1, 0}, {1, 0}, // left, right
	{-1, 1}, {0, 1}, {1, 1}, // down-left, down, down-right
}

func MakeImagePointSquareBool(max int) (mapping map[image.Point]bool) {
	mapping = make(map[image.Point]bool)
	for x := 0; x < max; x++ {
		for y := 0; y < max; y++ {
			mapping[image.Point{x, y}] = false
		}
	}
	return
}

func MakeImagePointMap(lines []string) (mapping map[image.Point]rune) {
	mapping = make(map[image.Point]rune)
	for y, s := range lines {
		for x, r := range s {
			mapping[image.Point{x, y}] = r
		}
	}
	return
}
func MakeImagePointMapRect(lines []string) (mapping map[image.Point]rune, rect image.Rectangle) {
	mapping = make(map[image.Point]rune)
	for y, s := range lines {
		for x, r := range s {
			mapping[image.Point{x, y}] = r
		}
	}
	rect = image.Rect(0, 0, len(lines[0])-1, len(lines)-1)
	return
}

func MakeIntImagePointMap(lines []string) (mapping map[image.Point]int, rect image.Rectangle) {
	mapping = MakeIntImagePoint(lines)
	rect = image.Rect(0, 0, len(lines[0])-1, len(lines)-1)
	return
}

func MakeIntImagePoint(lines []string) (mapping map[image.Point]int) {
	mapping = make(map[image.Point]int)
	for y, s := range lines {
		for x, r := range strings.Split(s, "") {
			mapping[image.Point{x, y}] = Atoi(r)
		}
	}
	return
}

func Adj(p, d image.Point) image.Point {
	return p.Add(d)
}
