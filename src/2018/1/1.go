package main

import (
	"fmt"
	"go_advent/src/2018/1/input"
	"strconv"
	"strings"
)

const (
	ToString      = "%q"
	SingleQuotes  = "'"
	BlankString   = ""
	CommaAndSpace = ", "
	NewLine       = "\n"
	DoubleQuotes  = "\""
)

// Atoi is an awful name. Just awful
var strToInt = strconv.Atoi

func convertStringPoints(str string) (string, int) {
	var direction = fmt.Sprintf(ToString, str[0])
	var amount = str[1:]

	var cleanedDirection = strings.ReplaceAll(direction, SingleQuotes, BlankString)
	var cleanedAmount = strings.ReplaceAll(amount, DoubleQuotes, BlankString)

	var amountAsInt, _ = strToInt(cleanedAmount)
	return cleanedDirection, amountAsInt
}

// second part
func main() {
	var operations = strings.Split(input.Input5, NewLine)
	var operationsLength = len(operations)

	var frequencyCounter = map[int]int{}

	var isFreqFoundTwice bool
	var frequency, iterator int
	for !isFreqFoundTwice {
		var currentOperation = operations[iterator]
		var direction, amount = convertStringPoints(currentOperation)

		if direction == "+" {
			frequency += amount
		} else {
			frequency -= amount
		}

		frequencyCounter[frequency]++

		if frequencyCounter[frequency] > 1 {
			isFreqFoundTwice = true
		} else {
			iterator++
			if iterator >= operationsLength {
				iterator = 0
			}
		}
	}
	fmt.Println(frequency)
}

// first part
// func main() {
// 	var split = strings.Split(input.Input5, NewLine)

// 	var frequency int
// 	for _, change := range split {
// 		var direction, amount = convertStringPoints(change)
// 		if direction == "+" {
// 			frequency += amount
// 		} else {
// 			frequency -= amount
// 		}
// 	}
// 	fmt.Println(frequency)
// }
