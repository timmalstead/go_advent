package main

import (
	"fmt"
	"go_advent/src/2015/2/input"
	"math"
	"strconv"
	"strings"
)

func main() {
	var splitInput = strings.Split(input.PuzzleInput, "\n")

	var totalWrappingPaper int
	for _, input := range splitInput {
		var dimensions = strings.Split(input, "x")

		var length, lErr = strconv.Atoi(dimensions[0])
		var width, wErr = strconv.Atoi(dimensions[1])
		var height, hErr = strconv.Atoi(dimensions[2])

		if lErr == nil && wErr == nil && hErr == nil {
			var sidesForPaper = []int{2 * length * width, 2 * width * height, +2 * height * length}
			var totalArea, smallestSide = 0, math.MaxInt

			for _, sideArea := range sidesForPaper {
				totalArea += sideArea

				if sideArea < smallestSide {
					smallestSide = sideArea
				}
			}

			totalWrappingPaper += totalArea + (smallestSide / 2)
		}
	}

	fmt.Println(totalWrappingPaper)
}
