package main

import (
	"fmt"
	"go_advent/src/2025/6/input"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var splitInput = strings.Split(input.TestInput, "\n")
	var twoDSlice = [][]any{}

	var lineLength int

	for _, line := range splitInput {
		var lineSlice = []any{}

		var trimLine = strings.Trim(line, " ")
		var splitLine = strings.Fields(trimLine)

		for _, numOrOperator := range splitLine {
			var trimChar = strings.TrimFunc(numOrOperator, unicode.IsSpace)

			var maybeConvertedInt, err = strconv.Atoi(trimChar)

			var convertedIntOrOperator any
			if err != nil {
				convertedIntOrOperator = numOrOperator
			} else {
				convertedIntOrOperator = maybeConvertedInt
			}

			lineSlice = append(lineSlice, convertedIntOrOperator)
		}
		twoDSlice = append(twoDSlice, lineSlice)
		lineLength = len(lineSlice)
	}

	var rotatedSlice = [][]any{}

	for y, subSlice := range twoDSlice {
		var rotatedArrayToAdd = []any{}
		for x, intOrOperator := range subSlice {
			if x == 0 {
				rotatedArrayToAdd = append(rotatedArrayToAdd, []any{})
			}
			rotatedArrayToAdd[x] = append(rotatedArrayToAdd[x], twoDSlice[y][x])
		}
	}

	fmt.Println(twoDSlice, lineLength)
}
