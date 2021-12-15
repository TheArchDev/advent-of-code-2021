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

func getInsertionRules(inputData []string) map[string]string {
	insertionRules := make(map[string]string)
	for i := 0; i < len(inputData); i++ {
		insertionRule := strings.Split(inputData[i], "->")
		insertionRules[strings.TrimSpace(insertionRule[0])] = strings.TrimSpace(insertionRule[1])
	}
	return insertionRules
}

func processInput(inputData []string) (template string, insertionRules map[string]string) {
	template = inputData[0]
	insertionRules = getInsertionRules(inputData[2:])
	return
}

func calculateNextStep(inputPolymer string, insertionRules map[string]string) (outputPolymer string) {
	// for i := 0; i < len(inputPolymer)-1; i++{

	// }
	outputPolymer = inputPolymer
	return
}

func calculateFinalPolymer(polymer string, insertionRules map[string]string, numberOfSteps int) string {
	for i := 0; i < numberOfSteps; i++ {
		polymer = calculateNextStep(polymer, insertionRules)
		fmt.Println(i, polymer)
	}
	return polymer
}

func main() {
	inputData := getInputData()
	template, insertionRules := processInput(inputData)
	fmt.Println(template)
	fmt.Println(insertionRules)

	numberOfSteps := 2
	finalPolymer := calculateFinalPolymer(template, insertionRules, numberOfSteps)
	fmt.Println(finalPolymer)
	fmt.Println(len(finalPolymer))
}
