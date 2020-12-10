package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var inputFile = flag.String("inputFile", "input.txt", "The relative filepath to the problem input")

func partOne() {
	adapters, _ := parseInput()
	differences := make(map[int]int)
	start := 0
	visit(start, adapters, differences)
	fmt.Println(differences[1] * differences[3])
}

func partTwo() {
	adapters, maxVal := parseInput()
	start := 0
	results := make(map[int]int)
	fmt.Println(numPaths(start, adapters, maxVal, results))
}

func numPaths(currAdapter int, adapters map[int]bool, maxVal int, res map[int]int) int {
	if currAdapter == maxVal {
		return 1
	}
	if _, cached := res[currAdapter]; cached {
		return res[currAdapter]
	}

	pathCount := 0
	for i := 1; i <= 3; i++ {
		neighbor := currAdapter + i
		if _, available := adapters[neighbor]; available {
			pathCount += numPaths(neighbor, adapters, maxVal, res)
		}
	}

	res[currAdapter] = pathCount
	return res[currAdapter]
}

func visit(currAdapter int, adapters map[int]bool, differences map[int]int) {
	for i := 1; i <= 3; i++ {
		neighbor := currAdapter + i
		if _, available := adapters[neighbor]; available {
			differences[i]++
			visit(neighbor, adapters, differences)
			break
		}
	}
}

func parseInput() (map[int]bool, int) {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal("Failed to open the problem input")
	}

	scanner := bufio.NewScanner(f)
	adapters := make(map[int]bool)
	maxJoltage := 0
	for scanner.Scan() {
		jolt, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Failed to convert string to int")
		}
		if jolt > maxJoltage {
			maxJoltage = jolt
		}
		adapters[jolt] = true
	}

	maxJoltage += 3
	adapters[maxJoltage] = true
	return adapters, maxJoltage
}

func main() {
	partOne()
	partTwo()
}
