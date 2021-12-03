package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part_one(instructions [][]string) int {
	var depth int
	var position int
	for _, instruction := range instructions {
		units, _ := strconv.Atoi(instruction[1])
		switch instruction[0] {
		case "forward":
			position += units
		case "up":
			depth -= units
		case "down":
			depth += units
		}
	}
	return depth * position
}

func part_two(instructions [][]string) int {
	var depth int
	var position int
	var aim int
	for _, instruction := range instructions {
		units, _ := strconv.Atoi(instruction[1])
		switch instruction[0] {
		case "forward":
			position += units
			depth += aim * units
		case "up":
			aim -= units
		case "down":
			aim += units
		}
	}
	return depth * position
}

func main() {
	filePtr := flag.String("file", "input.txt", "input file location")
	flag.Parse()

	file, err := os.Open(*filePtr)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var instructions [][]string
	for scanner.Scan() {
		instruction := strings.Fields(scanner.Text())
		instructions = append(instructions, instruction)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(part_one(instructions))
	fmt.Println(part_two(instructions))
}
