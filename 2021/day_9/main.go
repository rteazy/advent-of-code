package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var inputFile = flag.String("inputFile", "input.txt", "The relative filepath to the problem input")

type Coordinate struct {
	x, y int
}

func partOne() {
	grid := parseInput()
	risk := make([]int, 0)
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			neighbors := []Coordinate{
				{i - 1, j},
				{i + 1, j},
				{i, j - 1},
				{i, j + 1},
			}
			adjacentCount := 0
			numValid := 0
			for _, neighbor := range neighbors {
				x, y := neighbor.x, neighbor.y
				if x < m && x >= 0 && y < n && y >= 0 {
					adjacentCount++
					if grid[i][j] < grid[x][y] {
						numValid++
					}
				}
			}

			if numValid == adjacentCount {
				risk = append(risk, grid[i][j]+1)
			}

		}
	}

	total := 0
	for _, val := range risk {
		total += val
	}
	fmt.Printf("Part One: %d\n", total)
}

func partTwo() {
	grid := parseInput()
	m, n := len(grid), len(grid[0])
	basins := getBasins(grid)
	lengths := make([]int, 0)

	for _, basin := range basins {
		path := make(map[Coordinate]bool)
		path[basin] = true
		queue := []Coordinate{basin}
		for len(queue) > 0 {
			curr := queue[0]
			i, j := curr.x, curr.y
			queue = queue[1:]

			neighbors := []Coordinate{
				{i - 1, j},
				{i + 1, j},
				{i, j - 1},
				{i, j + 1},
			}

			for _, neighbor := range neighbors {
				x, y := neighbor.x, neighbor.y
				if _, visited := path[neighbor]; !visited &&
					x >= 0 && x < m && y >= 0 && y < n &&
					grid[x][y] > grid[curr.x][curr.y] && grid[x][y] != 9 {
					path[neighbor] = true
					queue = append(queue, neighbor)
				}
			}
		}
		lengths = append(lengths, len(path))
	}

	//  record length of path
	sort.Ints(lengths)
	n = len(lengths)
	results := lengths[n-1] * lengths[n-2] * lengths[n-3]
	fmt.Printf("Part 2: %d\n", results)
}

func getBasins(grid [][]int) []Coordinate {
	basins := make([]Coordinate, 0)
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			neighbors := make([][]int, 0)
			up := []int{i - 1, j}
			down := []int{i + 1, j}
			left := []int{i, j - 1}
			right := []int{i, j + 1}
			neighbors = append(neighbors, up)
			neighbors = append(neighbors, down)
			neighbors = append(neighbors, left)
			neighbors = append(neighbors, right)
			adjacentCount := 0
			numValid := 0
			for _, neighbor := range neighbors {
				x, y := neighbor[0], neighbor[1]
				if x < m && x >= 0 && y < n && y >= 0 {
					adjacentCount++
					if grid[i][j] < grid[x][y] {
						numValid++
					}
				}
			}

			if numValid == adjacentCount {
				basins = append(basins, Coordinate{i, j})
			}
		}
	}

	return basins
}

func parseInput() [][]int {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	grid := make([][]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		nums := make([]int, 0)
		for _, c := range line {
			num, _ := strconv.Atoi(string(c))
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
