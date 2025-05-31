package main

import (
	"fmt"
	"go_advent/src/2020/1/input"
	"strconv"
	"strings"
)

// first part
func main() {
	var reportEntries = strings.Split(input.PuzzleInput, "\n")

	var entriesFound bool
	var firstMultiplier, secondMultiplier int
	for i, firstEntry := range reportEntries {
		var firstToInt, err = strconv.Atoi(firstEntry)
		if err == nil {
			for j := i + 1; j < len(reportEntries); j++ {
				var entryToCompare = reportEntries[j]
				var compareAsInt, err = strconv.Atoi(entryToCompare)
				if err == nil {
					if firstToInt+compareAsInt == 2020 {
						firstMultiplier = firstToInt
						secondMultiplier = compareAsInt
						entriesFound = true
						break
					}
				}
			}
			if entriesFound {
				break
			}
		} else {
			fmt.Println(err)
		}
	}
	fmt.Println(firstMultiplier * secondMultiplier)
}
