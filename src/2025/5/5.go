package main

import (
	"fmt"
	"go_advent/src/2025/5/input"
	"strconv"
	"strings"
)

type Range struct {
	lowerBound, upperBound int
}

func main() {
	var splitInput = strings.Split(input.PuzzleInput, "\n\n")
	var freshIngredientRanges = strings.Split(splitInput[0], "\n")

	var freshIngredientIdRanges = []Range{}

	for _, ingredientRange := range freshIngredientRanges {
		var splitRange = strings.Split(ingredientRange, "-")

		var lowerBound, _ = strconv.Atoi(splitRange[0])
		var upperBound, _ = strconv.Atoi(splitRange[1])

		freshIngredientIdRanges = append(freshIngredientIdRanges, Range{lowerBound, upperBound})
	}

	var totalOfFreshIngredients int

	var ingredientList = strings.Split(splitInput[1], "\n")
	for _, ingredientId := range ingredientList {
		var idToCheck, _ = strconv.Atoi(ingredientId)

		for _, freshIngredientRange := range freshIngredientIdRanges {
			if idToCheck >= freshIngredientRange.lowerBound && idToCheck <= freshIngredientRange.upperBound {
				totalOfFreshIngredients++
				break
			}
		}
	}

	fmt.Println(totalOfFreshIngredients)
}
