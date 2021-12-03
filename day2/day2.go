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

	var depth int
	var position int
	for _, instruction := range instructions {
		steps, _ := strconv.Atoi(instruction[1])
		switch instruction[0] {
		case "forward":
			position += steps
		case "up":
			depth -= steps
		case "down":
			depth += steps
		}
	}
	fmt.Println(depth * position)
}
