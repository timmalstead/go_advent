package main

import (
	"fmt"
	"go_advent/src/2025/3/input"
	"strconv"
	"strings"
)

func main() {
	var splitInput = strings.Split(input.PuzzleInput, "\n")
	var total int

	for _, input := range splitInput {
		var inputLen = len(input)

		var firstLargestNumber = "0"
		var startSecondSearchIndex int

		for i, num := range input[:inputLen-1] {
			var runeToString = string(num)
			if runeToString > firstLargestNumber {
				firstLargestNumber = runeToString
				startSecondSearchIndex = i + 1
			}
		}

		var nextLargestNumber = "0"
		for i := startSecondSearchIndex; i < inputLen; i++ {
			var runeToString = string(input[i])
			if runeToString > nextLargestNumber {
				nextLargestNumber = runeToString
			}
		}

		var sum, _ = strconv.Atoi(firstLargestNumber + nextLargestNumber)
		total += sum
	}

	fmt.Println(total)
}
