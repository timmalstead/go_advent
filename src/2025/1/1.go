package main

import (
	"fmt"
	"go_advent/src/2025/1/input"
	"strconv"
	"strings"
)

func numIsTooHigh(numToTest int) bool {
	return numToTest > 99
}

func numIsTooLow(numToTest int) bool {
	return numToTest < 0
}

func numIsOutOfBounds(numToTest int) bool {
	return numIsTooHigh(numToTest) || numIsTooLow(numToTest)
}

func recurseCircle(resolvedNumber int) int {
	var result int

	if numIsTooHigh(resolvedNumber) {
		result = resolvedNumber - 100
	}
	if numIsTooLow(resolvedNumber) {
		result = resolvedNumber + 100
	}

	if numIsOutOfBounds(result) {
		result = recurseCircle(result)
	}

	return result
}

func main() {
	var splitInput = strings.Split(input.FirstPuzzleInput, "\n")
	var currentPosition, zeroPositionTally = 50, 0

	for _, input := range splitInput {
		var direction = input[0:1]
		var amountToMove, err = strconv.Atoi(input[1:])

		if err == nil {

			if direction == "L" {
				currentPosition -= amountToMove
			} else {
				currentPosition += amountToMove
			}

			if numIsOutOfBounds(currentPosition) {
				currentPosition = recurseCircle(currentPosition)
			}

			if currentPosition == 0 {
				zeroPositionTally++
			}

		} else {
			fmt.Println(err)
		}
	}

	fmt.Println(zeroPositionTally)
}
