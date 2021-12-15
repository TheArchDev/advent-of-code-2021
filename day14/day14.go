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

func getRuleCount(template string) map[string]int {
	ruleCount := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		ruleCount[template[i:i+2]]++
	}
	return ruleCount
}

func getRuleMapping(insertionRules map[string]string) map[string][2]string {
	ruleMapping := make(map[string][2]string)
	for key, value := range insertionRules {
		ruleMapping[key] = [2]string{string(key[0]) + value, value + string(key[1])}
	}
	return ruleMapping
}

func getFinalRuleCount(ruleCount map[string]int, ruleMapping map[string][2]string, numberOfSteps int) map[string]int {
	// TODO consider starting at step 1, to align with problem statement
	for i := 0; i < numberOfSteps; i++ {
		ruleCount = getNextRuleCount(ruleCount, ruleMapping)
	}
	return ruleCount
}

func getNextRuleCount(inputRulecount map[string]int, ruleMapping map[string][2]string) map[string]int {
	ruleCount := make(map[string]int)
	for rule, count := range inputRulecount {
		for _, mappedRule := range ruleMapping[rule] {
			ruleCount[mappedRule] += count
		}
	}
	return ruleCount
}

func getElementCount(ruleCount map[string]int, template string) map[string]int {
	elementCount := make(map[string]int)
	for rule, count := range ruleCount {
		elementCount[string(rule[0])] += count
		elementCount[string(rule[1])] += count
	}
	for element, count := range elementCount {
		elementCount[element] = count / 2
	}
	// TODO tidy up
	elementCount[string(template[0])]++
	elementCount[string(template[len(template)-1])]++
	return elementCount
}

func processElementCounts(elementCount map[string]int) (mostCommonCount, leastCommonCount int) {
	mostCommonCount = 0
	// TODO calculate leastCommonCount properly! Currently just take N as the minimum as know it'll be there
	leastCommonCount = elementCount["N"]
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

	ruleCount := getRuleCount(template)
	ruleMapping := getRuleMapping(insertionRules)

	numberOfSteps := 40
	finalRuleCount := getFinalRuleCount(ruleCount, ruleMapping, numberOfSteps)

	elementCount := getElementCount(finalRuleCount, template)

	mostCommonCount, leastCommonCount := processElementCounts(elementCount)
	fmt.Println(partOneResult(mostCommonCount, leastCommonCount))
}
