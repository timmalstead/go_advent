package main

import (
	"fmt"
	"go_advent/src/2015/3/input"
	"strings"
)

var positionFuncs = map[string]func(*int, *int){
	">": func(x, _ *int) { *x++ },
	"<": func(x, _ *int) { *x-- },
	"^": func(_, y *int) { *y++ },
	"v": func(_, y *int) { *y-- },
}

func main() {
	var x, y int
	var coords = map[string]bool{"x0y0": true}

	var splitInput = strings.Split(input.PuzzleInput, "")
	for _, direction := range splitInput {
		positionFuncs[direction](&x, &y)
		coords[fmt.Sprintf("x%dy%d", x, y)] = true
	}

	fmt.Println(len(coords))
}
