package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	filePtr := flag.String("file", "input.txt", "input file location")
	flag.Parse()

	file, _ := os.Open(*filePtr)
	scanner := bufio.NewScanner(file)
	var allLines [][2][2]int
	var horizontalLines [][2][2]int
	var verticalLines [][2][2]int
	var largestX int
	var largestY int
	for scanner.Scan() {
		input := strings.Split(strings.Join(strings.Fields(scanner.Text()), ""), "->")

		var initial [2]int
		for index, value := range strings.Split(input[0], ",") {
			initial[index], _ = strconv.Atoi(value)
		}
		if initial[0] > largestX {
			largestX = initial[0]
		}
		if initial[1] > largestY {
			largestY = initial[1]
		}

		var final [2]int
		for index, value := range strings.Split(input[1], ",") {
			final[index], _ = strconv.Atoi(value)
		}
		if final[0] > largestX {
			largestX = final[0]
		}
		if final[1] > largestY {
			largestY = final[1]
		}

		element := [2][2]int{initial, final}
		allLines = append(allLines, element)
		if element[0][1] == element[1][1] {
			horizontalLines = append(horizontalLines, element)
		}
		// duplicate!!
		if element[0][0] == element[1][0] {
			verticalLines = append(verticalLines, element)
		}
	}
	// fmt.Println(allLines)
	fmt.Println("verticalLines", verticalLines)
	fmt.Println()

	// generate a two dimensional matrix of counts
	counts := make([][]int, largestY+1)
	for index := range counts {
		counts[index] = make([]int, largestX+1)
	}

	// loop through all sets of coordinates
	// increment the relevant spots in the 2d matrix
	// VERTICAL
	for _, lines := range verticalLines {
		fmt.Println("lines", lines)
		yCoordinates := []int{lines[0][1], lines[1][1]}
		sort.Ints(yCoordinates)
		fmt.Println("yCoordinates", yCoordinates)
		for yCoordinate := yCoordinates[0]; yCoordinate <= yCoordinates[1]; yCoordinate++ {
			xCoordinate := lines[0][0]
			fmt.Println("stepping here", xCoordinate, yCoordinate)
			counts[yCoordinate][xCoordinate]++
		}
		fmt.Println()
	}

	// HORIZONTAL
	fmt.Println("horizontalLines", horizontalLines)
	for _, lines := range horizontalLines {
		fmt.Println("lines", lines)
		xCoordinates := []int{lines[0][0], lines[1][0]}
		sort.Ints(xCoordinates)
		fmt.Println("xCoordinates", xCoordinates)
		for xCoordinate := xCoordinates[0]; xCoordinate <= xCoordinates[1]; xCoordinate++ {
			yCoordinate := lines[0][1]
			fmt.Println("stepping here", xCoordinate, yCoordinate)
			counts[yCoordinate][xCoordinate]++
		}
		fmt.Println()
	}

	var total int
	fmt.Println(counts)
	for _, row := range counts {
		for _, count := range row {
			if count >= 2 {
				total++
			}
		}
	}
	fmt.Println(total)

}
