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

func process_diagnostics(diagnostics []string, column int)(processed_diagnostics []string) {
	fmt.Println("incoming diagnostics", diagnostics)
	most_common_bits, _ := gamma_and_epsilon(diagnostics)
	fmt.Println("most_common_bits", most_common_bits)
	for _, diagnostic := range diagnostics {
		if diagnostic[column] == most_common_bits[column] {
			processed_diagnostics = append(processed_diagnostics, diagnostic)
		}
	}
	// fmt.Println(processed_diagnostics)
	return	
}

func process_co2_scrubber(diagnostics []string, column int)(processed_diagnostics []string) {
	fmt.Println("incoming diagnostics", diagnostics)
	_, least_common_bits := gamma_and_epsilon(diagnostics)
	fmt.Println("least_common_bits", least_common_bits)
	for _, diagnostic := range diagnostics {
		if diagnostic[column] == least_common_bits[column] {
			processed_diagnostics = append(processed_diagnostics, diagnostic)
		}
	}
	// fmt.Println(processed_diagnostics)
	return	
}

func part_two(diagnostics []string) int64 {
	fmt.Println("initial diagnostics", diagnostics)
	processed_diagnostics := process_diagnostics(diagnostics, 0)
	fmt.Println("first process", processed_diagnostics)

	// processed_diagnostics = process_diagnostics(processed_diagnostics, 1)
	// fmt.Println("second process", processed_diagnostics)	
	
	for i := 1; len(processed_diagnostics) > 1; i++{
		fmt.Println(i)
		processed_diagnostics = process_diagnostics(processed_diagnostics, i)	
	}
	fmt.Println(processed_diagnostics)



	processed_co2 := process_co2_scrubber(diagnostics, 0)
	for i := 1; len(processed_co2) > 1; i++{
		fmt.Println(i)
		processed_co2 = process_co2_scrubber(processed_co2, i)	
	}
	fmt.Println("processed_co2", processed_co2)


	oxygen, _ := strconv.ParseInt(processed_diagnostics[0], 2, 64)
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
