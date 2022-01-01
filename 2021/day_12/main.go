package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var inputFile = flag.String("inputFile", "input.txt", "The relative filepath to the problem input")

func partOne() {
	graph := parseInput()
	path := make(map[string]bool)
	path["start"] = true
	res := visit(graph, "start", path)
	fmt.Printf("Part 1: %d\n", res)
}

func partTwo() {
	graph := parseInput()
	visited := make(map[string]bool)
	visited["start"] = true
	path := make([]string, 0)
	path = append(path, "start")
	result := make(map[string]bool)
	visitAdjusted(graph, "start", visited, false, path, result)
	fmt.Printf("Part 2: %d\n", len(result))
}

func visitAdjusted(graph map[string][]string, node string, visited map[string]bool, visitedTwice bool, currPath []string, pathsFound map[string]bool) int {
	if node == "end" {
		s := strings.Join(currPath, ",")
		pathsFound[s] = true
		return 1
	}

	total := 0
	for _, neighbor := range graph[node] {
		if _, explored := visited[neighbor]; !explored {
			upperCase := false
			if neighbor[0] == strings.ToUpper(neighbor)[0] {
				upperCase = true
			}
			currPath = append(currPath, neighbor)
			if !upperCase {
				if visitedTwice {
					visited[neighbor] = true
					total += visitAdjusted(graph, neighbor, visited, visitedTwice, currPath, pathsFound)
					delete(visited, neighbor)
				} else {
					visited[neighbor] = true
					total += visitAdjusted(graph, neighbor, visited, false, currPath, pathsFound)
					delete(visited, neighbor)
					total += visitAdjusted(graph, neighbor, visited, true, currPath, pathsFound)
				}
			} else {
				total += visitAdjusted(graph, neighbor, visited, visitedTwice, currPath, pathsFound)
			}
			currPath = currPath[:len(currPath)-1]
		}
	}

	return total
}

func visit(graph map[string][]string, node string, path map[string]bool) int {
	if node == "end" {
		return 1
	}

	total := 0
	for _, neighbor := range graph[node] {
		if _, visited := path[neighbor]; !visited {
			upperCase := false
			if neighbor[0] == strings.ToUpper(neighbor)[0] {
				upperCase = true
			}
			if !upperCase {
				path[neighbor] = true
			}
			total += visit(graph, neighbor, path)
			if !upperCase {
				delete(path, neighbor)
			}
		}
	}

	return total
}

func parseInput() map[string][]string {
	graph := make(map[string][]string)
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, "-")
		a, b := tokens[0], tokens[1]
		if _, exists := graph[a]; !exists {
			graph[a] = make([]string, 0)
		}
		graph[a] = append(graph[a], b)
		if _, exists := graph[b]; !exists {
			graph[b] = make([]string, 0)
		}
		graph[b] = append(graph[b], a)
	}
	return graph
}

func main() {
	partOne()
	partTwo()
}
