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

var inputFile = flag.String("inputFile", "input.txt", "The relative filepath to the problem input")

func partOne() {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	scanner := bufio.NewScanner(f)
	states := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		numStr := strings.Split(line, ",")
		for _, num := range numStr {
			val, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("Failed on %s\n", num)
			}
			states = append(states, val)
		}
	}

	for i := 0; i < 80; i++ {
		states = nextDay(states)
	}
	fmt.Printf("Part One: %d\n", len(states))
}

func partTwo() {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	scanner := bufio.NewScanner(f)
	states := make([]int, 9)
	for scanner.Scan() {
		line := scanner.Text()
		numStr := strings.Split(line, ",")
		for _, num := range numStr {
			val, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("Failed on %s\n", num)
			}
			states[val]++
		}
	}

	for i := 0; i < 256; i++ {
		states = nextDayUsingCounts(states)
	}
	total := 0
	for _, val := range states {
		total += val
	}
	fmt.Printf("Part 2: %d\n", total)

}

func nextDayUsingCounts(counts []int) []int {
	nextDay := make([]int, 9)
	for i := 1; i <= 8; i++ {
		nextDay[i-1] = counts[i]
	}
	if counts[0] > 0 {
		nextDay[8] += counts[0]
		nextDay[6] += counts[0]
	}

	return nextDay
}

func nextDay(states []int) []int {
	nextStates := make([]int, len(states))
	for i, v := range states {
		if v == 0 {
			nextStates = append(nextStates, 8)
			nextStates[i] = 6
		} else {
			nextStates[i] = v - 1
		}
	}
	return nextStates
}

func main() {
	partOne()
	partTwo()
}
