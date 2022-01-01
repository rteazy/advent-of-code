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

type Coordinate struct {
	x, y int
}

func partOne() {
	grid := parseInput()
	total := 0
	numSteps := 100
	for i := 0; i < numSteps; i++ {
		total += step(grid)
	}
	fmt.Printf("Part 1: %d\n", total)
}

func partTwo() {
	grid := parseInput()
	step := 0
	numSteps := 1000000
	for i := 0; i < numSteps; i++ {
		step++
		if synchronized(grid) {
			break
		}
	}
	fmt.Printf("Part 1: %d\n", step)
}

func step(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	flashCount := 0
	stack := make([]Coordinate, 0)
	newFlash := 10
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] += 1
			if grid[i][j] == newFlash {
				stack = append(stack, Coordinate{i, j})
			}
		}
	}
	for len(stack) > 0 {
		curr := stack[0]
		x, y := curr.x, curr.y
		neighbors := []Coordinate{
			{x - 1, y},
			{x + 1, y},
			{x, y - 1},
			{x, y + 1},
			{x - 1, y - 1},
			{x - 1, y + 1},
			{x + 1, y - 1},
			{x + 1, y + 1},
		}
		for _, neighbor := range neighbors {
			i, j := neighbor.x, neighbor.y
			if i >= 0 && i < m && j >= 0 && j < n {
				grid[i][j]++
				if grid[i][j] == newFlash {
					stack = append(stack, Coordinate{i, j})
				}
			}
		}
		stack = stack[1:]
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] >= newFlash {
				grid[i][j] = 0
				flashCount++
			}
		}
	}
	return flashCount
}

func synchronized(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	flashCount := 0
	stack := make([]Coordinate, 0)
	newFlash := 10
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] += 1
			if grid[i][j] == newFlash {
				stack = append(stack, Coordinate{i, j})
			}
		}
	}
	for len(stack) > 0 {
		curr := stack[0]
		x, y := curr.x, curr.y
		neighbors := []Coordinate{
			{x - 1, y},
			{x + 1, y},
			{x, y - 1},
			{x, y + 1},
			{x - 1, y - 1},
			{x - 1, y + 1},
			{x + 1, y - 1},
			{x + 1, y + 1},
		}
		for _, neighbor := range neighbors {
			i, j := neighbor.x, neighbor.y
			if i >= 0 && i < m && j >= 0 && j < n {
				grid[i][j]++
				if grid[i][j] == newFlash {
					stack = append(stack, Coordinate{i, j})
				}
			}
		}
		stack = stack[1:]
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] >= newFlash {
				grid[i][j] = 0
				flashCount++
			}
		}
	}
	return flashCount == m*n
}

func parseInput() [][]int {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	scanner := bufio.NewScanner(f)
	grid := make([][]int, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		nums := make([]int, 0)
		for _, v := range line {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("Failed on %s\n", v)
			}
			nums = append(nums, num)
		}
		grid = append(grid, nums)
	}
	return grid
}

func main() {
	partOne()
	partTwo()
}
