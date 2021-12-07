package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func part_one(initialPositions []int, smallestPosition, largestPosition int) int {
	smallestTotalSteps := math.Inf(1)
	for position := smallestPosition; position <= largestPosition; position++ {
		var totalSteps float64
		for _, crabPosition := range initialPositions {
			totalSteps += math.Abs(float64(crabPosition - position))
		}
		if totalSteps < smallestTotalSteps {
			smallestTotalSteps = totalSteps
		}
	}
	return int(smallestTotalSteps)
}

func part_two(initialPositions []int, smallestPosition, largestPosition int) int {
	smallestTotalFuel := math.Inf(1)
	for position := smallestPosition; position <= largestPosition; position++ {
		var totalFuel float64
		for _, crabPosition := range initialPositions {
			positionDifference := math.Abs(float64(crabPosition - position))
			totalFuel += ((positionDifference + 1) * positionDifference) / 2
		}
		if totalFuel < smallestTotalFuel {
			smallestTotalFuel = totalFuel
		}
	}
	return int(smallestTotalFuel)
}

func main() {
	filePtr := flag.String("file", "input.txt", "input file location")
	flag.Parse()

	contents, _ := os.ReadFile(*filePtr)
	inputs := strings.Split(string(contents), ",")

	initialPositions := make([]int, len(inputs))
	smallestPosition, _ := strconv.Atoi(inputs[0])
	var largestPosition int
	for index, input := range inputs {
		inputInt, _ := strconv.Atoi(input)
		if inputInt > largestPosition {
			largestPosition = inputInt
		} else if inputInt < smallestPosition {
			smallestPosition = inputInt
		}
		initialPositions[index] = inputInt
	}

	fmt.Println(part_one(initialPositions, smallestPosition, largestPosition))
	fmt.Println(part_two(initialPositions, smallestPosition, largestPosition))
}
