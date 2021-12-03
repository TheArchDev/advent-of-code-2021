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

func part_one(depths []int) (count int) {
	previous_depth := depths[0]
	for _, depth := range depths {
		if depth > previous_depth {
			count++
		}
		previous_depth = depth
	}
	return count
}

func main() {
	filePtr := flag.String("file", "input.txt", "input file location")
	flag.Parse()

	file, err := os.Open(*filePtr)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var depths []int
	for scanner.Scan() {
		depth, _ := strconv.Atoi(strings.ReplaceAll(scanner.Text(), " ", ""))
		depths = append(depths, depth)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(part_one(depths))
}
