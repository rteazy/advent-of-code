package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

var inputFile = flag.String("inputFile", "input.txt", "The relative filepath to the problem input")

func partOne() {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal("Failed to open the problem input file")
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
	}

}

func main() {
	partOne()
}
