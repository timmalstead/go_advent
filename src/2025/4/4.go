package main

import (
	"fmt"
	"go_advent/src/2025/4/input"
	"strings"
	"sync"
)

type Roll struct {
	x, y  int
	value string
}

func formatCoord(x, y int) string {
	return fmt.Sprintf("x%dy%d", x, y)
}

func main() {
	var coords = make(map[string]Roll)
	var totalNumberOfAccessibleRolls int
	var mutex sync.Mutex

	// i believe for this part it's simplest to keep it syncronous
	var splitInput = strings.Split(input.PuzzleInput, "\n")
	for y, line := range splitInput {
		var splitLine = strings.Split(line, "")
		for x, char := range splitLine {
			if char == "@" {
				coords[formatCoord(x, y)] = Roll{x: x, y: y, value: char}
			}
		}
	}

	var mainWaitGroup sync.WaitGroup
	var readSurroundingCoords = func(key string) {
		defer mainWaitGroup.Done()
		var totalNumberOfSurroundingRolls int

		var coord = coords[key]
		var x, y = coord.x, coord.y

		var coordsToCheck = []string{
			formatCoord(x+1, y),   // right
			formatCoord(x+1, y+1), // right down
			formatCoord(x, y+1),   // down
			formatCoord(x-1, y+1), // left down
			formatCoord(x-1, y),   // left
			formatCoord(x-1, y-1), // left up
			formatCoord(x, y-1),   // up
			formatCoord(x+1, y-1), // up right
		}

		var subWaitGroup sync.WaitGroup
		var checkSingleCoord = func(coordKey string) {
			defer subWaitGroup.Done()

			var _, checkingCoordExists = coords[coordKey]
			if checkingCoordExists {
				mutex.Lock()
				totalNumberOfSurroundingRolls++
				mutex.Unlock()
			}
		}

		subWaitGroup.Add(len(coordsToCheck))
		for _, coordKey := range coordsToCheck {
			go checkSingleCoord(coordKey)
		}

		subWaitGroup.Wait()
		if totalNumberOfSurroundingRolls < 4 {
			// the mutex is important here as I am accessing and modifying a shared variable
			mutex.Lock()
			totalNumberOfAccessibleRolls++
			mutex.Unlock()
		}
	}

	mainWaitGroup.Add(len(coords))
	for key := range coords {
		go readSurroundingCoords(key)
	}
	mainWaitGroup.Wait()

	fmt.Println(totalNumberOfAccessibleRolls)
}
