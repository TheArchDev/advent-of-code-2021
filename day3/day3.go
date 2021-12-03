package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func gamma_and_epsilon(diagnostics []string) (gamma_binary, epsilon_binary string) {
	column_counts := make([]int, len(diagnostics[0]))
	for _, diagnostic := range diagnostics {
		for index, value := range diagnostic {
			if value == '1' {
				column_counts[index]++
			} else {
				column_counts[index]--
			}
		}
	}

	for _, count := range column_counts {
		if count >= 0 {
			gamma_binary += "1"
			epsilon_binary += "0"
		} else {
			gamma_binary += "0"
			epsilon_binary += "1"
		}
	}
	return
}

func part_one(diagnostics []string) int64 {
	gamma_binary, epsilon_binary := gamma_and_epsilon(diagnostics)
	gamma, _ := strconv.ParseInt(gamma_binary, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilon_binary, 2, 64)
	return gamma * epsilon
}

func process_oxygen(diagnostics []string, column int) (processed_diagnostics []string) {
	most_common_bits, _ := gamma_and_epsilon(diagnostics)
	for _, diagnostic := range diagnostics {
		if diagnostic[column] == most_common_bits[column] {
			processed_diagnostics = append(processed_diagnostics, diagnostic)
		}
	}
	return
}

func process_co2_scrubber(diagnostics []string, column int) (processed_diagnostics []string) {
	_, least_common_bits := gamma_and_epsilon(diagnostics)
	for _, diagnostic := range diagnostics {
		if diagnostic[column] == least_common_bits[column] {
			processed_diagnostics = append(processed_diagnostics, diagnostic)
		}
	}
	return
}

func part_two(diagnostics []string) int64 {
	processed_oxygen := process_oxygen(diagnostics, 0)
	for i := 1; len(processed_oxygen) > 1; i++ {
		processed_oxygen = process_oxygen(processed_oxygen, i)
	}

	processed_co2 := process_co2_scrubber(diagnostics, 0)
	for i := 1; len(processed_co2) > 1; i++ {
		processed_co2 = process_co2_scrubber(processed_co2, i)
	}

	oxygen, _ := strconv.ParseInt(processed_oxygen[0], 2, 64)
	co2, _ := strconv.ParseInt(processed_co2[0], 2, 64)
	return oxygen * co2
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
	fmt.Println(part_two(diagnostics))
}
