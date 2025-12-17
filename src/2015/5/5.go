package main

import (
	"fmt"
	"go_advent/src/2015/5/input"
	"strings"
	"sync"
)

type VowelCounter struct {
	indice int
	char   string
}

var vowels = map[string]int{"a": 0, "e": 0, "i": 0, "o": 0, "u": 0}
var disallowedStrings = map[string]int{"ab": 0, "cd": 0, "pq": 0, "xy": 0}

func main() {
	var splitInput = strings.Split(input.PuzzleInput, "\n")
	var niceStringsCount int

	var mutex sync.Mutex
	var waitGroup sync.WaitGroup

	var testString = func(str string) {
		defer waitGroup.Done()

		var vowelCount = map[int]int{}
		var hasDoubleLetter bool
		var hasDisallowedLetters bool

		var countVowels = func(counter VowelCounter) {
			var _, charIsVowel = vowels[counter.char]
			if charIsVowel {
				vowelCount[counter.indice]++
			}
		}

		for i := 1; i < len(str); i++ {
			var firstLetter = string(str[i-1])
			var secondLetter = string(str[i])

			var compoundLetters = fmt.Sprintf("%s%s", firstLetter, secondLetter)

			var _, isDisallowedString = disallowedStrings[compoundLetters]

			if isDisallowedString {
				hasDisallowedLetters = true
				break
			}

			countVowels(VowelCounter{indice: i - 1, char: firstLetter})
			countVowels(VowelCounter{indice: i, char: secondLetter})

			if firstLetter == secondLetter {
				hasDoubleLetter = true
			}

		}

		var isNiceString = len(vowelCount) >= 3 && hasDoubleLetter && !hasDisallowedLetters

		if isNiceString {
			mutex.Lock()
			niceStringsCount++
			mutex.Unlock()
		}
	}

	waitGroup.Add(len(splitInput))
	for _, inputLine := range splitInput {
		go testString(inputLine)
	}
	waitGroup.Wait()

	fmt.Println(niceStringsCount)
}
