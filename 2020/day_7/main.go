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

var inputFile = flag.String("inputFile", "input.txt", "File path for problem input.")
var numberOfPathsToNode = 0

// Bag represents the inner bag
type Bag struct {
	count    int
	bagColor string
}

func partOne() {
	graph := buildAdjacencyList()
	containsBag := make(map[string]bool)
	for node := range graph {
		if _, visited := containsBag[node]; !visited {
			findBag(graph, node, containsBag)
		}
	}

	fmt.Println(numberOfPathsToNode)
}

func findBag(graph map[string][]Bag, node string, containsBag map[string]bool) bool {
	if node == "shiny gold" {
		return true
	}

	if found, exists := containsBag[node]; exists {
		return found
	}

	foundTarget := false
	for _, neighbor := range graph[node] {
		if findBag(graph, neighbor.bagColor, containsBag) {
			foundTarget = true
			numberOfPathsToNode++
			break
		}
	}

	containsBag[node] = foundTarget
	return containsBag[node]
}

func partTwo() {
	graph := buildAdjacencyList()
	fmt.Println(countBags(graph, "shiny gold") - 1)
}

func countBags(graph map[string][]Bag, color string) int {
	bagCount := 1
	for _, neighbor := range graph[color] {
		res := neighbor.count * countBags(graph, neighbor.bagColor)
		bagCount += res
	}
	return bagCount
}

func buildAdjacencyList() map[string][]Bag {
	flag.Parse()
	f, _ := os.Open(*inputFile)

	scanner := bufio.NewScanner(f)
	graph := make(map[string][]Bag)

	for scanner.Scan() {
		line := scanner.Text()
		bags := strings.Split(line, " contain ")
		outerBag, contains := bags[0], bags[1]
		outer := strings.Split(outerBag, " ")
		outerBagColor := strings.Join(outer[:len(outer)-1], " ")

		contains = strings.TrimRight(contains, ".")
		innerBags := strings.Split(contains, ", ")
		for _, bag := range innerBags {
			inner := strings.Split(bag, " ")
			if inner[0] == "no" {
				graph[outerBagColor] = []Bag{}
				continue
			}
			innerBagCount, err := strconv.Atoi(inner[0])
			if err != nil {
				log.Fatal("Failed to parse the inner bag count")
			}
			innerBagColor := strings.Join(inner[1:len(inner)-1], " ")
			bag := Bag{innerBagCount, innerBagColor}
			graph[outerBagColor] = append(graph[outerBagColor], bag)
		}
	}

	return graph
}

func main() {
	partOne()
	partTwo()
}
