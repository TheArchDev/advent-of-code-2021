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

	scanner.Scan()
	previous_depth, _ := strconv.Atoi(strings.ReplaceAll(scanner.Text(), " ", ""))
	var count int
	for scanner.Scan() {
		depth, _ := strconv.Atoi(strings.ReplaceAll(scanner.Text(), " ", ""))
		if depth > previous_depth {
			count++
		}
		previous_depth = depth
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
