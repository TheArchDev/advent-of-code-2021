package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePtr := flag.String("file", "input.txt", "input file location")
	flag.Parse()

	file, _ := os.Open(*filePtr)
	scanner := bufio.NewScanner(file)

	var inputs []string
	for scanner.Scan(){
		inputs = append(inputs, scanner.Text())
	}

	numberOfRows := len(inputs)
	grids := make([][]int, numberOfRows)
	numberOfColumns := len(inputs[0])
	for index := range grids {
		grids[index] = make([]int, numberOfColumns)
	}
	for j, input := range inputs {
		for i, heightString := range strings.Split(input, "") {
			heightInt, _ := strconv.Atoi(heightString)
			grids[j][i] = heightInt
		}
	}
	var countOfLowValues int
	for j, row := range grids {
		for i, height := range row {
			fmt.Println()
			var heightsToCompare []int
			abovePosition := j - 1
			if abovePosition >= 0 {
				heightsToCompare = append(heightsToCompare, grids[abovePosition][i])
			}
			belowPosition := j + 1
			if belowPosition < numberOfRows {
				heightsToCompare = append(heightsToCompare, grids[belowPosition][i])
			}
			leftPosition := i - 1
			if leftPosition >= 0 {
				heightsToCompare = append(heightsToCompare, grids[j][leftPosition])
			}
			rightPosition := i + 1
			if rightPosition < numberOfColumns {
				heightsToCompare = append(heightsToCompare, grids[j][rightPosition])
			}
			fmt.Println("heightsToCompare", heightsToCompare)
			lowestHeight := true
			for _, valueToCompare := range heightsToCompare {
				if valueToCompare <= height {
					lowestHeight = false
				}
			}
			if lowestHeight {
				fmt.Println("location", j, i)
				fmt.Println("height", height)
				countOfLowValues += height + 1			}
		}
	}
	fmt.Println(countOfLowValues)

}