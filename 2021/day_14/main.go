package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

var inputFile = flag.String("inputFile", "input.txt", "The relative filepath to the problem input")

type Pair struct {
	a, b string
	key  string
}

func NewPair(a, b string) Pair {
	pair := Pair{a, b, a + b}
	return pair
}

func partOne() {
	fmt.Printf("Part 1: %d\n", run(10))
}

func partTwo() {
	fmt.Printf("Part 2: %d\n", run(40))
}

func run(numSteps int) int {
	template, rules := parseInput()
	counts := make(map[string]int)
	for i := 0; i < len(template); i++ {
		counts[string(template[i])]++
	}

	pairs := initPairs(template)
	for i := 0; i < numSteps; i++ {
		pairs = nextPairs(rules, counts, pairs)
	}

	res := calculateDifference(counts)
	return res
}

func nextPairs(rules map[string]string, characterCount map[string]int, pairCounts map[Pair]int) map[Pair]int {
	newPairs := make(map[Pair]int)
	for pair, count := range pairCounts {
		newChar := rules[pair.key]
		characterCount[newChar] += count
		newPairA, newPairB := NewPair(pair.a, newChar), NewPair(newChar, pair.b)
		newPairs[newPairA] += count
		newPairs[newPairB] += count
	}

	return newPairs
}

func parseInput() (string, map[string]string) {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	template := scanner.Text()
	scanner.Scan()
	rules := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " -> ")
		pair, val := tokens[0], tokens[1]
		rules[pair] = val
	}

	return template, rules
}

func initPairs(template string) map[Pair]int {
	pairs := make(map[Pair]int, 0)
	for i := 0; i < len(template)-1; i++ {
		a, b := string(template[i]), string(template[i+1])
		pairs[NewPair(a, b)]++
	}

	return pairs
}

func calculateDifference(counts map[string]int) int {
	minFreq, maxFreq := math.MaxInt64, 0
	for _, count := range counts {
		if count < minFreq {
			minFreq = count
		}
		if count > maxFreq {
			maxFreq = count
		}
	}

	return maxFreq - minFreq
}

func main() {
	partOne()
	partTwo()
}
