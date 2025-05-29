package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	DoubleQuotes = "\""
	SingleQuotes = "'"
	BlankString  = ""
	Space        = " "
	ToString     = "%q"
	NewLines     = `\n`
	Tabs         = `\t`
)

func convertCharToInt(char string) (int, error) {
	var cleanedChar = strings.ReplaceAll(char, SingleQuotes, BlankString)

	return strconv.Atoi(cleanedChar)
}

func descendingOrder(a, b int) int {
	return b - a
}

// second part
func main() {
	var data, err = os.ReadFile("./src/2017/2/input_2.txt")

	if err == nil {

		var text = fmt.Sprintf(ToString, data)
		var splitOnNewLines = strings.Split(text, NewLines)

		var checkSum int
		for _, line := range splitOnNewLines {
			var trimmedLine = strings.Trim(line, DoubleQuotes)
			var splitOnSpaces = strings.Split(trimmedLine, Tabs)

			var ints = []int{}
			for _, char := range splitOnSpaces {
				var convertedInt, err = convertCharToInt(char)
				if err == nil {
					ints = append(ints, convertedInt)
				}
			}

			slices.SortFunc(ints, descendingOrder)

			for i, firstInteger := range ints {
				for j := i + 1; j < len(ints); j++ {
					var secondInteger = ints[j]

					var doesDivideEvenly = firstInteger%secondInteger == 0
					if doesDivideEvenly {
						checkSum += firstInteger / secondInteger
						break
					}

				}
			}
		}
		fmt.Println(checkSum)

	} else {

		fmt.Println(err)

	}
}

// first part
// func main() {
// 	var data, err = os.ReadFile("./src/2017/2/input_2.txt")

// 	if err == nil {

// 		var text = fmt.Sprintf(ToString, data)
// 		var splitOnNewLines = strings.Split(text, NewLines)

// 		var checkSum int
// 		for _, line := range splitOnNewLines {
// 			var trimmedLine = strings.Trim(line, DoubleQuotes)
// 			var splitOnSpaces = strings.Split(trimmedLine, Tabs)

// 			var smallest, largest = math.MaxInt, -1
// 			for _, char := range splitOnSpaces {
// 				var convertedInt, err = convertCharToInt(char)

// 				if err == nil {
// 					if convertedInt < smallest {
// 						smallest = convertedInt
// 					}

// 					if convertedInt > largest {
// 						largest = convertedInt
// 					}
// 				} else {
// 					fmt.Println(err)
// 				}
// 			}
// 			checkSum += (largest - smallest)
// 		}

// 		fmt.Println(checkSum)

// 	} else {

// 		fmt.Println(err)

// 	}
// }
