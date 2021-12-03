package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filePtr := flag.String("file", "input.txt", "input file location")
	flag.Parse()

	file, err := os.Open(*filePtr)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var text []string
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		text = append(text, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(text)
}
