package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
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
	positions := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		posStrs := strings.Split(line, ",")
		for _, posStr := range posStrs {
			val, err := strconv.Atoi(posStr)
			if err != nil {
				log.Fatalf("Failed for: %s\n", posStr)
			}
			positions = append(positions, val)
		}
	}

	smallest := math.MaxInt32
	for _, a := range positions {
		cost := 0
		for _, b := range positions {
			diff := int(math.Abs(float64(a - b)))
			cost += diff
		}
		if cost < smallest {
			smallest = cost
		}
	}

	fmt.Printf("Part 1: %d\n", smallest)
}

func partTwo() {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	scanner := bufio.NewScanner(f)
	positions := make([]int, 0)
	start, end := math.MaxInt32, math.MinInt32
	for scanner.Scan() {
		line := scanner.Text()
		posStrs := strings.Split(line, ",")
		for _, posStr := range posStrs {
			val, err := strconv.Atoi(posStr)
			if err != nil {
				log.Fatalf("Failed for: %s\n", posStr)
			}
			if val < start {
				start = val
			}
			if val > end {
				end = val
			}
			positions = append(positions, val)
		}
	}

	// Keep track of the fuel costs in fuelCosts map[int][int]
	maxDifference := end - start
	fuelCosts := make(map[int]int)
	initialCost := 0
	for i := 0; i <= maxDifference; i++ {
		initialCost = initialCost + i
		fuelCosts[i] = initialCost
	}

	// from start->end, calculate cost at each pos
	// if fuelCost[abs(pos - curr)] < smallest, set fuel cost
	smallest := math.MaxInt32
	for curr := start; curr <= end; curr++ {
		total := 0
		for _, pos := range positions {
			diff := int(math.Abs(float64(pos - curr)))
			cost, exists := fuelCosts[diff]
			if !exists {
				log.Fatalf("Could not find difference in fuel costs: %d\n", diff)
			}
			total += cost
		}
		if total < smallest {
			smallest = total
		}
	}

	fmt.Printf("Part 2: %d\n", smallest)
}

func main() {
	partOne()
	partTwo()
}
