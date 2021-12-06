package main

import (
	// "bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePtr := flag.String("file", "input.txt", "input file location")
	daysPtr := flag.Int("days", 80, "number of days to simulate")
	flag.Parse()

	fileContent, err := os.ReadFile(*filePtr)
	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(string(fileContent), ",")
	timers := make([]int, len(inputs))
	for index, timer := range inputs {
		timers[index], _ = strconv.Atoi(timer)
	}

	for i := 1; i <= *daysPtr; i++ {
		var newTimers []int
		for index, timer := range timers {
			if timer == 0 {
				timer = 6
				newTimers = append(newTimers, 8)
			} else {
				timer--
			}
			timers[index] = timer
		}
		timers = append(timers, newTimers...)
	}
	fmt.Println("Number of lanterns", len(timers))
}
