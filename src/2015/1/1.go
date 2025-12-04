package main

import (
	"fmt"
	"go_advent/src/2015/1/input"
	"strings"
)

func main() {
	var splitInput = strings.Split(input.PuzzleInput, "")
	var floor, firstBasementIndex int

	for i, brace := range splitInput {
		if brace == "(" {
			floor++
		} else {
			floor--
		}

		if floor == -1 && firstBasementIndex == 0 {
			firstBasementIndex = i + 1
		}
	}

	fmt.Println(floor, firstBasementIndex)
}
