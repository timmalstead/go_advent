package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

const (
	DoubleQuotes = "\""
	Space        = " "
	ToString     = "%q"
	NewLines     = `\n`
	SingleQuotes = "'"
	BlankString  = ""
)

// second part
func main() {
	var data, err = os.ReadFile("./src/2017/4/input_2.txt")

	if err == nil {
		var text = fmt.Sprintf(ToString, data)
		var splitOnNewLines = strings.Split(text, NewLines)

		var numberOfValidPassPhrases int
		for _, line := range splitOnNewLines {
			var trimmedLine = strings.Trim(line, DoubleQuotes)
			var splitOnSpaces = strings.Split(trimmedLine, Space)

			var counterHolder = []map[string]int{}
			for _, word := range splitOnSpaces {
				var charCounter = make(map[string]int)
				for _, char := range word {
					var runeConvertedToString = fmt.Sprintf(ToString, char)
					var cleanedChar = strings.ReplaceAll(runeConvertedToString, SingleQuotes, BlankString)
					charCounter[cleanedChar]++
				}
				counterHolder = append(counterHolder, charCounter)
			}

			var allWordsAreUnique = true
			for i, currentCounter := range counterHolder {
				for j := i + 1; j < len(counterHolder); j++ {
					var nextCounter = counterHolder[j]

					var wordsAreAnagrams = reflect.DeepEqual(currentCounter, nextCounter)
					if wordsAreAnagrams {
						allWordsAreUnique = false
						break
					}
				}
				if !allWordsAreUnique {
					break
				}
			}

			if allWordsAreUnique {
				numberOfValidPassPhrases++
			}
		}
		fmt.Println(numberOfValidPassPhrases)
	} else {

		fmt.Println(err)

	}
}

// // first part
// func main() {
// 	var data, err = os.ReadFile("./src/2017/4/input_2.txt")

// 	if err == nil {
// 		var text = fmt.Sprintf(ToString, data)
// 		var splitOnNewLines = strings.Split(text, NewLines)

// 		var numberOfValidPassPhrases int
// 		for _, line := range splitOnNewLines {
// 			var trimmedLine = strings.Trim(line, DoubleQuotes)
// 			var splitOnSpaces = strings.Split(trimmedLine, Space)

// 			var wordFreqency, allWordsAreUnique = make(map[string]int), true
// 			for _, word := range splitOnSpaces {
// 				wordFreqency[word]++
// 				if wordFreqency[word] > 1 {
// 					allWordsAreUnique = false
// 					break
// 				}
// 			}

// 			if allWordsAreUnique {
// 				numberOfValidPassPhrases++
// 			}
// 		}
// 		fmt.Println(numberOfValidPassPhrases)
// 	} else {

// 		fmt.Println(err)

// 	}
// }
