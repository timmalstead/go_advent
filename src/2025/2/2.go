package main

import (
	"fmt"
	"go_advent/src/2025/2/input"
	"strconv"
	"strings"
)

func main() {
	var splitInput = strings.Split(input.PuzzleInput, ",")

	var total int
	for _, input := range splitInput {
		var bounds = strings.Split(input, "-")

		var lowerBound, _ = strconv.Atoi(bounds[0])
		var upperBound, _ = strconv.Atoi(bounds[1])

		for i := lowerBound; i <= upperBound; i++ {
			var intToStr = strconv.Itoa(i)
			var strLen = len(intToStr)
			var isEven = strLen%2 == 0

			if !isEven {
				continue
			}

			var halfLen = strLen / 2
			var firstHalf, secondHalf = intToStr[0:halfLen], intToStr[halfLen:]

			if firstHalf == secondHalf {
				total += i
			}
		}
	}

	fmt.Println(total)
}
