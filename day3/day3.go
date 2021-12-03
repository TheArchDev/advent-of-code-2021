package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part_one(diagnostics []string) int64 {
	column_counts := make([]int, len(diagnostics[1]))
	for _, diagnostic := range diagnostics {
		for index, value := range diagnostic {
			if value == '1' {
				column_counts[index]++
			} else {
				column_counts[index]--
			}
		}
	}

	var gamma_string string
	var epsilon_string string
	for _, count := range column_counts {
		if count < 0 {
			gamma_string += "0"
			epsilon_string += "1"
		} else {
			gamma_string += "1"
			epsilon_string += "0"
		}
	}
	gamma, _ := strconv.ParseInt(gamma_string, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilon_string, 2, 64)
	return gamma * epsilon
}

func main() {
	filePtr := flag.String("file", "input.txt", "input file location")
	flag.Parse()

	file, err := os.Open(*filePtr)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var diagnostics []string
	for scanner.Scan() {
		diagnostic := scanner.Text()
		diagnostics = append(diagnostics, diagnostic)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(part_one(diagnostics))
}
