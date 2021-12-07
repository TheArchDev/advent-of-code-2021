package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type fn func(float64, float64) float64

func calculatePositionDifference(positionOne, positionTwo float64) (positionDifference float64) {
	positionDifference = math.Abs(positionTwo - positionOne)
	return
}

func calculateFuelDifference(positionOne, positionTwo float64) (fuelDifference float64) {
	positionDifference := calculatePositionDifference(positionOne, positionTwo)
	fuelDifference = ((positionDifference + 1) * positionDifference) / 2
	return
}

func calculate(initialPositions []int, f fn, smallestPosition, largestPosition int) int {
	smallestTotal := math.Inf(1)
	for position := smallestPosition; position <= largestPosition; position++ {
		var total float64
		for _, initialPosition := range initialPositions {
			total += f(float64(initialPosition), float64(position))
		}
		if total < smallestTotal {
			smallestTotal = total
		}
	}
	return int(smallestTotal)
}

func partOne(initialPositions []int, smallestPosition, largestPosition int) int {
	return calculate(initialPositions, calculatePositionDifference, smallestPosition, largestPosition)
}

func partTwo(initialPositions []int, smallestPosition, largestPosition int) int {
	return calculate(initialPositions, calculateFuelDifference, smallestPosition, largestPosition)
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

	fmt.Println(partOne(initialPositions, smallestPosition, largestPosition))
	fmt.Println(partTwo(initialPositions, smallestPosition, largestPosition))
}
