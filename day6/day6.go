package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calculate_timers(timerCounts map[int]int) map[int]int {
	updatedCounts := make(map[int]int)
	for i := 0; i < 8; i++ {
		updatedCounts[i] = timerCounts[i+1]
	}
	updatedCounts[6] += timerCounts[0]
	updatedCounts[8] = timerCounts[0]

	return updatedCounts
}

func main() {
	filePtr := flag.String("file", "input.txt", "input file location")
	daysPtr := flag.Int("days", 80, "number of days to simulate")
	flag.Parse()

	fileContent, err := os.ReadFile(*filePtr)
	if err != nil {
		log.Fatal(err)
	}
	inputs := strings.Split(string(fileContent), ",")

	initialTimers := make([]int, len(inputs))
	for index, timer := range inputs {
		initialTimers[index], _ = strconv.Atoi(timer)
	}

	timerCounts := make(map[int]int)
	for _, timerValue := range initialTimers {
		timerCounts[timerValue]++
	}

	for i := 1; i <= *daysPtr; i++ {
		timerCounts = calculate_timers(timerCounts)
	}

	var total int
	for _, count := range timerCounts {
		total += count
	}
	fmt.Println(total)
}
