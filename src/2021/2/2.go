package main

import (
	"fmt"
	"go_advent/src/2021/2/input"
	"strconv"
	"strings"
	"sync"
)

var positionFuncs = map[string]func(*int, *int, int){
	"forward": func(x, _ *int, movement int) { *x += movement },
	"up":      func(_, y *int, movement int) { *y -= movement },
	"down":    func(_, y *int, movement int) { *y += movement },
}

func main() {
	var splitInput = strings.Split(input.PuzzleInput, "\n")
	var x, y int

	var waitGroup sync.WaitGroup
	var mutex sync.Mutex

	var calculatePosition = func(line string) {
		defer waitGroup.Done()

		var splitLine = strings.Split(line, " ")

		var direction = splitLine[0]
		var movement, err = strconv.Atoi(splitLine[1])

		if err == nil {
			mutex.Lock()
			positionFuncs[direction](&x, &y, movement)
			mutex.Unlock()
		}
	}

	waitGroup.Add(len(splitInput))
	for _, line := range splitInput {
		go calculatePosition(line)
	}
	waitGroup.Wait()

	fmt.Println(x * y)
}
