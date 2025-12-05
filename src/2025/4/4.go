package main

import (
	"fmt"
	"go_advent/src/2025/4/input"
	"strings"
	"sync"
)

type Roll struct {
	x     int
	y     int
	value string
}

const CoordTemplate = "x%dy%d"

func main() {
	var coords = make(map[string]Roll)
	var totalNumberOfAccessibleRolls int
	var mutex sync.Mutex

	// i believe for this part it's simplest to keep it syncronous
	var splitInput = strings.Split(input.PuzzleInput, "\n")
	for y, line := range splitInput {
		var splitLine = strings.Split(line, "")
		for x, char := range splitLine {
			coords[fmt.Sprintf(CoordTemplate, x, y)] = Roll{x: x, y: y, value: char}
		}
	}

	var mainWaitGroup sync.WaitGroup
	var readSurroundingCoords = func(key string) {
		defer mainWaitGroup.Done()

		var coord = coords[key]
		var x, y, currentValue = coord.x, coord.y, coord.value

		if currentValue == "@" {
			var totalNumberOfSurroundingRolls int

			var coordsToCheck = []string{
				fmt.Sprintf(CoordTemplate, x+1, y),   // right
				fmt.Sprintf(CoordTemplate, x+1, y+1), // right down
				fmt.Sprintf(CoordTemplate, x, y+1),   // down
				fmt.Sprintf(CoordTemplate, x-1, y+1), // left down
				fmt.Sprintf(CoordTemplate, x-1, y),   // left
				fmt.Sprintf(CoordTemplate, x-1, y-1), // left up
				fmt.Sprintf(CoordTemplate, x, y-1),   // up
				fmt.Sprintf(CoordTemplate, x+1, y-1), // up right
			}

			for _, coordKey := range coordsToCheck {
				var checkingCoord, checkingCoordExists = coords[coordKey]

				if checkingCoordExists && checkingCoord.value == "@" {
					totalNumberOfSurroundingRolls++
				}
			}

			if totalNumberOfSurroundingRolls < 4 {
				// the mutex is important here as I am accessing and modifying a shared variable
				mutex.Lock()
				totalNumberOfAccessibleRolls++
				mutex.Unlock()
			}
		}
	}

	mainWaitGroup.Add(len(coords))

	for key := range coords {
		go readSurroundingCoords(key)
	}

	mainWaitGroup.Wait()

	fmt.Println(totalNumberOfAccessibleRolls)
}
