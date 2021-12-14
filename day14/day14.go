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

func getInsertionRules(inputData []string) [][2]string {
	insertionRules := make([][2]string, len(inputData))
	for i := 0; i < len(inputData); i++ {
		var insertionRule [2]string
		for index, value := range strings.Split(inputData[i], "->") {
			insertionRule[index] = strings.TrimSpace(value)
		}
		insertionRules[i] = insertionRule
	}
	return insertionRules
}

func processInput(inputData []string) (template string, insertionRules [][2]string) {
	template = inputData[0]
	insertionRules = getInsertionRules(inputData[2:])
	return
}

func main() {
	inputData := getInputData()
	template, insertionRules := processInput(inputData)
	fmt.Println(template)
	fmt.Println(insertionRules)
	fmt.Println(len(insertionRules))
}
