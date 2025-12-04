package main

import (
	"fmt"
	"go_advent/src/2015/3/input"
	"strings"
)

func main() {
	var splitInput = strings.Split(input.PuzzleInput, "")
	var x, y int
	var coords = map[string]bool{"x0y0": true}

	for _, input := range splitInput {
		if input == ">" {
			x++
		}
		if input == "<" {
			x--
		}
		if input == "^" {
			y++
		}
		if input == "v" {
			y--
		}

		var templatedCoord = fmt.Sprintf("x%dy%d", x, y)
		coords[templatedCoord] = true
	}

	fmt.Println(len(coords))
}
