package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	DoubleQuotes = "\""
	SingleQuotes = "'"
	BlankString  = ""
	ToString     = "%q"
)

func convertCharToInt(char rune) (int, error) {
	var charToString = fmt.Sprintf(ToString, char)
	var cleanedChar = strings.ReplaceAll(charToString, SingleQuotes, BlankString)

	return strconv.Atoi(cleanedChar)
}

// second part
func main() {
	var data, err = os.ReadFile("./src/2017/1/input_2.txt")

	if err == nil {

		var text = fmt.Sprintf(ToString, data)
		var cleanedText = strings.ReplaceAll(text, DoubleQuotes, BlankString)

		var total int

		var totalLength = len(cleanedText)
		var loopEnd, lookAheadIndex = totalLength - 1, totalLength / 2

		for i, currentChar := range cleanedText {
			var iterator = i

			var nextCharIndex = iterator + lookAheadIndex
			var nextCharIndexIsOutOfBounds = nextCharIndex > loopEnd

			var nextChar rune
			if nextCharIndexIsOutOfBounds {
				nextChar = rune(cleanedText[nextCharIndex-totalLength])
			} else {
				nextChar = rune(cleanedText[nextCharIndex])
			}

			var currentNum, err1 = convertCharToInt(currentChar)
			var nextNum, err2 = convertCharToInt(nextChar)

			var noErrs = err1 == nil && err2 == nil
			if noErrs {

				if currentNum == nextNum {
					total += currentNum
				}

			} else {

				fmt.Println(err1, err2)

			}
		}

		fmt.Println(total)

	} else {

		fmt.Println(err)

	}
}

// first part
// func main() {
// 	var data, err = os.ReadFile("./src/2017/1/input_2.txt")

// 	if err == nil {

// 		var text = fmt.Sprintf(ToString, data)
// 		var cleanedText = strings.ReplaceAll(text, DoubleQuotes, BlankString)

// 		var total int
// 		var loopEnd = len(cleanedText) - 1

// 		for i, currentChar := range cleanedText {
// 			var iterator = i
// 			var isLastChar = iterator == loopEnd

// 			var nextChar rune
// 			if isLastChar {
// 				var firstChar = rune(cleanedText[0])
// 				nextChar = firstChar
// 			} else {
// 				var charPlusOne = rune(cleanedText[iterator+1])
// 				nextChar = charPlusOne
// 			}

// 			var currentNum, err1 = convertCharToInt(currentChar)
// 			var nextNum, err2 = convertCharToInt(nextChar)

// 			var noErrs = err1 == nil && err2 == nil
// 			if noErrs {

// 				if currentNum == nextNum {
// 					total += currentNum
// 				}

// 			} else {

// 				fmt.Println(err1, err2)

// 			}
// 		}

// 		fmt.Println(total)

// 	} else {

// 		fmt.Println(err)

// 	}
// }
