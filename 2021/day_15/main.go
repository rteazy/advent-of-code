package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input.txt", "The relative filepath to the problem input")

type Point struct {
	i, j int
}

type Node struct {
	i, j     int
	visited  bool
	distance int
}

func NewNode(i, j int) *Node {
	return &Node{
		i, j, false, math.MaxInt64,
	}
}

func partOne() {
	grid := parseInput()
	nodes := make(map[Point]*Node)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			nodes[Point{i, j}] = NewNode(i, j)
		}
	}
	fmt.Printf("Part 1: %d\n", minDist(grid, nodes))
}

func partTwo() {
	grid, nodes := GetFullGrid()
	fmt.Printf("Part 2: %d\n", minDist(grid, nodes))
}

func minDist(grid [][]int, nodes map[Point]*Node) int {
	startPoint := NewNode(0, 0)
	startPoint.distance = 0
	queue := []*Node{startPoint}
	m, n := len(grid), len(grid[0])

	var endPoint *Node
	for len(queue) > 0 {
		curr := queue[0]
		i, j := curr.i, curr.j
		if i == m-1 && j == n-1 {
			endPoint = curr
		}
		if visited := curr.visited; !visited {
			curr.visited = true
			// look up neighbor in nodes map
			neighborPoints := []Point{
				{i - 1, j},
				{i + 1, j},
				{i, j - 1},
				{i, j + 1},
			}
			neighbors := make([]*Node, 0)
			for _, point := range neighborPoints {
				x, y := point.i, point.j
				if neighborNode, exists := nodes[point]; exists && x >= 0 && x < m && y >= 0 && y < n {
					neighbors = append(neighbors, neighborNode)
				}
			}
			for _, neighbor := range neighbors {
				x, y := neighbor.i, neighbor.j
				val := curr.distance + grid[x][y]
				if val < neighbor.distance {
					neighbor.distance = val
				}
				if !neighbor.visited {
					queue = append(queue, neighbor)
				}
			}
		}

		queue = queue[1:]
		sort.SliceStable(queue, func(i, j int) bool {
			return queue[i].distance < queue[j].distance
		})
	}

	return endPoint.distance
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
		line := strings.Split(scanner.Text(), "")
		nums := make([]int, len(line))
		for i, c := range line {
			val, err := strconv.Atoi(c)
			if err != nil {
				log.Fatalf("Could not convert %s\n", c)
			}
			nums[i] = val
		}
		grid = append(grid, nums)
	}

	return grid
}

func GetFullGrid() ([][]int, map[Point]*Node) {
	tileGrid := parseInput()
	nodes := make(map[Point]*Node)
	offsetCol, offsetRow := len(tileGrid[0]), len(tileGrid)
	m, n := offsetRow*5, offsetCol*5

	fullGrid := make([][]int, m)
	for i := 0; i < m; i++ {
		fullGrid[i] = make([]int, n)
	}

	for i := 0; i < offsetRow; i++ {
		for j := 0; j < offsetCol; j++ {
			for x := i; x < m; x += offsetRow {
				for y := j; y < n; y += offsetCol {
					if x == i && y == j {
						fullGrid[x][y] = tileGrid[i][j]
						nodes[Point{x, y}] = NewNode(x, y)
					} else if y == j { // get value from above
						val := fullGrid[x-offsetRow][y]
						if val == 9 {
							val = 1
						} else {
							val++
						}
						fullGrid[x][y] = val
						nodes[Point{x, y}] = NewNode(x, y)
					} else { // get value from left
						val := fullGrid[x][y-offsetCol]
						if val == 9 {
							val = 1
						} else {
							val++
						}
						fullGrid[x][y] = val
						nodes[Point{x, y}] = NewNode(x, y)
					}
				}
			}
		}
	}

	return fullGrid, nodes
}

func main() {
	partOne()
	partTwo()
}
