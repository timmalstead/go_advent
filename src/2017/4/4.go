package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	DoubleQuotes = "\""
	Space        = " "
	ToString     = "%q"
	NewLines     = `\n`
)

// first part
func main() {
	var data, err = os.ReadFile("./src/2017/4/input_2.txt")

	if err == nil {

		var text = fmt.Sprintf(ToString, data)
		var splitOnNewLines = strings.Split(text, NewLines)

		var numberOfValidPassPhrases int
		for _, line := range splitOnNewLines {
			var trimmedLine = strings.Trim(line, DoubleQuotes)
			var splitOnSpaces = strings.Split(trimmedLine, Space)

			var wordFreqency = make(map[string]int)
			for _, word := range splitOnSpaces {
				wordFreqency[word]++
			}

			var allWordsAreUnique bool = true
			for _, frequency := range wordFreqency {
				if frequency > 1 {
					allWordsAreUnique = false
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
