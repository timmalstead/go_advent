package main

import (
	"fmt"
	"go_advent/src/2021/1/input"
	"math"
	"strconv"
	"strings"
)

func main() {
	var currentSum = math.MaxInt
	var timesIncreased int

	var splitInput = strings.Split(input.PuzzleInput, "\n")

	for _, line := range splitInput {
		var sumCopy = currentSum

		var convertedInt, err = strconv.Atoi(line)

		if err == nil {
			currentSum = convertedInt

			if currentSum > sumCopy {
				timesIncreased++
			}
		}
	}
	fmt.Println(timesIncreased)
}
