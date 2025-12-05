package main

import (
	"fmt"
	"go_advent/src/2025/4/input"
	"strings"
	"sync"
)

type Roll struct {
	x, y int
}

func formatCoord(x, y int) string {
	return fmt.Sprintf("x%dy%d", x, y)
}

func main() {
	var coords = map[string]Roll{}

	var splitInput = strings.Split(input.PuzzleInput, "\n")
	for y, line := range splitInput {
		// keep the mapping sync

		var splitLine = strings.Split(line, "")
		for x, char := range splitLine {
			if char == "@" {
				coords[formatCoord(x, y)] = Roll{x, y}
			}
		}
	}

	var totalNumberOfAccessibleRolls int
	var mutex sync.Mutex
	var waitGroup sync.WaitGroup

	var searchSurroundingCoords = func(roll Roll) {
		defer waitGroup.Done()
		var totalNumberOfSurroundingRolls int

		var x, y = roll.x, roll.y
		var coordsToCheck = [8]string{
			formatCoord(x+1, y),   // →
			formatCoord(x+1, y+1), // ↘
			formatCoord(x, y+1),   // ↓
			formatCoord(x-1, y+1), // ↙
			formatCoord(x-1, y),   // ←
			formatCoord(x-1, y-1), // ↖
			formatCoord(x, y-1),   // ↑
			formatCoord(x+1, y-1), // ↗
		}

		for _, coordKey := range coordsToCheck {
			// no need to create a goroutine to check each direction

			var _, coordExists = coords[coordKey]
			if coordExists {
				totalNumberOfSurroundingRolls++
			}
		}

		if totalNumberOfSurroundingRolls < 4 {
			// use mutex for modifying shared variable
			mutex.Lock()
			totalNumberOfAccessibleRolls++
			mutex.Unlock()
		}
	}

	waitGroup.Add(len(coords))
	for _, roll := range coords {
		go searchSurroundingCoords(roll)
	}
	waitGroup.Wait()

	fmt.Println(totalNumberOfAccessibleRolls)
}
