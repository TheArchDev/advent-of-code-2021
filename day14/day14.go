package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func getInputData() (inputData []string) {
	filePtr := flag.String("file", "input.txt", "input file location")
	flag.Parse()

	file, _ := os.Open(*filePtr)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputData = append(inputData, scanner.Text())
	}
	return
}

func processInput(inputData []string) (template string, insertionRules map[string]string) {
	template = inputData[0]
	insertionRules = getInsertionRules(inputData[2:])
	return
}

func getInsertionRules(inputData []string) map[string]string {
	insertionRules := make(map[string]string)
	for i := 0; i < len(inputData); i++ {
		insertionRule := strings.Split(inputData[i], "->")
		insertionRules[strings.TrimSpace(insertionRule[0])] = strings.TrimSpace(insertionRule[1])
	}
	return insertionRules
}

func calculateFinalPolymer(polymer string, insertionRules map[string]string, numberOfSteps int) string {
	for i := 0; i < numberOfSteps; i++ {
		polymer = calculateNextStep(polymer, insertionRules)
	}
	return polymer
}

func calculateNextStep(inputPolymer string, insertionRules map[string]string) string {
	outputPolymer := string(inputPolymer[0])
	for i := 0; i < len(inputPolymer)-1; i++ {
		outputPolymer += insertionRules[string(inputPolymer[i:i+2])]
		outputPolymer += string(inputPolymer[i+1])
	}
	return outputPolymer
}

func getElementCount(polymer string) map[string]int {
	elementCount := make(map[string]int)
	for _, node := range polymer {
		elementCount[string(node)]++
	}
	return elementCount
}

func processElementCounts(elementCount map[string]int, totalElementCount int) (mostCommonCount, leastCommonCount int) {
	mostCommonCount = 0
	leastCommonCount = totalElementCount
	for _, value := range elementCount {
		if value > mostCommonCount {
			mostCommonCount = value
		}
		if value < leastCommonCount {
			leastCommonCount = value
		}
	}
	return
}

func partOneResult(mostCommonCount, leastCommonCount int) int {
	return mostCommonCount - leastCommonCount
}

func main() {
	inputData := getInputData()
	template, insertionRules := processInput(inputData)

	numberOfSteps := 10
	finalPolymer := calculateFinalPolymer(template, insertionRules, numberOfSteps)

	elementCount := getElementCount(finalPolymer)

	mostCommonCount, leastCommonCount := processElementCounts(elementCount, len(finalPolymer))
	fmt.Println(partOneResult(mostCommonCount, leastCommonCount))
}
