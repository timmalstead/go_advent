package main

import (
	"fmt"
	"go_advent/src/2018/2/input"
	"strings"
)

// first part
func main() {
	var idSlices = strings.Split(input.PuzzleInput, "\n")

	var appearsTwiceTally, appearsThriceTally int
	for _, id := range idSlices {
		var charCounter = map[rune]int{}
		for _, char := range id {
			charCounter[char]++
		}

		var doesAppearTwice, doesAppearThrice bool
		for _, frequency := range charCounter {
			if doesAppearTwice && doesAppearThrice {
				break
			}

			if frequency == 2 {
				doesAppearTwice = true
			}
			if frequency == 3 {
				doesAppearThrice = true
			}
		}

		if doesAppearTwice {
			appearsTwiceTally++
		}

		if doesAppearThrice {
			appearsThriceTally++
		}
	}

	var checksum = appearsTwiceTally * appearsThriceTally
	fmt.Println(checksum)
}
