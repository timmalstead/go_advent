package main

import (
	"fmt"
	"go_advent/src/2021/3/input"
	"strconv"
	"strings"
	"sync"
)

var bitSwap = map[string]string{
	"1": "0",
	"0": "1",
}

func main() {
	var splitInput = strings.Split(input.PuzzleInput, "\n")
	var positiveBits = map[int]int{}

	var waitGroup sync.WaitGroup
	var mutex sync.Mutex

	var findPositiveBits = func(binaryStr string) {
		defer waitGroup.Done()

		var splitBinary = strings.Split(binaryStr, "")
		for i, bitStr := range splitBinary {
			if bitStr == "1" {
				mutex.Lock()
				positiveBits[i]++
				mutex.Unlock()
			}
		}
	}

	waitGroup.Add(len(splitInput))
	for _, binaryStr := range splitInput {
		go findPositiveBits(binaryStr)
	}
	waitGroup.Wait()

	var gammaRate string
	for i := 0; i < len(positiveBits); i++ {
		var currentBit = positiveBits[i]
		if currentBit < len(splitInput)/2 {
			gammaRate += "0"
		} else {
			gammaRate += "1"
		}
	}

	var epsilonRate string
	for _, gammaRateBit := range gammaRate {
		epsilonRate += bitSwap[string(gammaRateBit)]
	}

	var convertedGamma, gammaErr = strconv.ParseInt(gammaRate, 2, 64)
	var convertedEpsilon, epsilonErr = strconv.ParseInt(epsilonRate, 2, 64)

	if gammaErr == nil && epsilonErr == nil {
		fmt.Println(convertedGamma * convertedEpsilon)
	}
}
