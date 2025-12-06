package main

import (
	"fmt"
	"go_advent/src/2025/5/input"
	"strconv"
	"strings"
	"sync"
)

type Range struct {
	lowerBound, upperBound int
}

func main() {
	var splitInput = strings.Split(input.PuzzleInput, "\n\n")
	var freshIngredientRanges = strings.Split(splitInput[0], "\n")

	var waitGroup sync.WaitGroup
	var mutex sync.Mutex

	var freshIngredientIdRanges = []Range{}

	var addRange = func(ingredientRange string) {
		defer waitGroup.Done()

		var splitRange = strings.Split(ingredientRange, "-")

		var lowerBound, _ = strconv.Atoi(splitRange[0])
		var upperBound, _ = strconv.Atoi(splitRange[1])

		mutex.Lock()
		freshIngredientIdRanges = append(freshIngredientIdRanges, Range{lowerBound, upperBound})
		mutex.Unlock()
	}

	waitGroup.Add(len(freshIngredientRanges))
	for _, ingredientRange := range freshIngredientRanges {
		go addRange(ingredientRange)
	}

	waitGroup.Wait()

	var totalOfFreshIngredients int
	var ingredientList = strings.Split(splitInput[1], "\n")

	var checkRange = func(ingredientId string) {
		defer waitGroup.Done()
		var idToCheck, _ = strconv.Atoi(ingredientId)

		for _, freshIngredientRange := range freshIngredientIdRanges {
			if idToCheck >= freshIngredientRange.lowerBound && idToCheck <= freshIngredientRange.upperBound {
				mutex.Lock()
				totalOfFreshIngredients++
				mutex.Unlock()
				break
			}
		}
	}

	waitGroup.Add(len(ingredientList))
	for _, ingredientId := range ingredientList {
		go checkRange(ingredientId)
	}

	waitGroup.Wait()

	fmt.Println(totalOfFreshIngredients)
}
