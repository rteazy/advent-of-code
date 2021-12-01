package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var inputFile = flag.String("inputFile", "input.txt", "The relative filepath to the problem input")

func partOne() {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("failed to open file")
	}

	scanner := bufio.NewScanner(f)
	prev := int(math.Inf(-1))
	count := -1
	for scanner.Scan() {
		line := scanner.Text()
		curr, _ := strconv.Atoi(line)
		if curr > prev {
			count++
		}
		prev = curr
	}

	fmt.Printf("Part 1: %d\n", count)
}

func partTwo() {
	//Sliding window of size 3
	// total, a, b, curr
	// if total > total - a + curr then inc count
	// total = total - a + curr

	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("failed to open file")
	}

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())
	total := int(math.Inf(-1))
	count := -1

	for scanner.Scan() {
		curr, _ := strconv.Atoi(scanner.Text())
		if total < a+b+curr {
			count++
		}
		total = a + b + curr
		a, b = b, curr
	}

	fmt.Printf("Part 2: %d\n", count)
}

func main() {
	partOne()
	partTwo()
}
