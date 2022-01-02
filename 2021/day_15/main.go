package main

import (
	"bufio"
	"container/heap"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
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
	cost     int
	distance float64
}

type MinHeap []*Node

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return float64(h[i].cost)+h[i].distance < float64(h[j].cost)+h[j].distance
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x interface{}) {
	item := x.(*Node)
	*h = append(*h, item)
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return item
}

func NewNode(i int, j int, distance float64) *Node {
	return &Node{
		i, j, false, math.MaxInt64, distance,
	}
}

func partOne() {
	grid := parseInput()
	nodes := make(map[Point]*Node)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			dist := euclideanDistance(grid, i, j)
			nodes[Point{i, j}] = NewNode(i, j, dist)
		}
	}
	fmt.Printf("Part 1: %d\n", minDist(grid, nodes))
}

func partTwo() {
	grid, nodes := GetFullGrid()
	fmt.Printf("Part 2: %d\n", minDist(grid, nodes))
}

func minDist(grid [][]int, nodes map[Point]*Node) int {
	startPoint := NewNode(0, 0, 0)
	startPoint.cost = 0
	queue := make(MinHeap, 0)
	heap.Init(&queue)
	heap.Push(&queue, startPoint)
	m, n := len(grid), len(grid[0])

	var endPoint *Node
	for len(queue) > 0 {
		curr := heap.Pop(&queue).(*Node)
		i, j := curr.i, curr.j
		if i == m-1 && j == n-1 {
			endPoint = curr
			break
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
				if neighborNode, exists := nodes[point]; exists && x >= 0 && x < m && y >= 0 && y < n && !neighborNode.visited {
					neighbors = append(neighbors, neighborNode)
				}
			}
			for _, neighbor := range neighbors {
				x, y := neighbor.i, neighbor.j
				val := curr.cost + grid[x][y]
				if val < neighbor.cost {
					neighbor.cost = val
				}
				if !neighbor.visited {
					heap.Push(&queue, neighbor)
				}
			}
		}
	}

	return endPoint.cost
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
					dist := euclideanDistance(fullGrid, x, y)
					if x == i && y == j {
						fullGrid[x][y] = tileGrid[i][j]
						nodes[Point{x, y}] = NewNode(x, y, dist)
					} else if y == j { // get value from above
						val := fullGrid[x-offsetRow][y]
						if val == 9 {
							val = 1
						} else {
							val++
						}
						fullGrid[x][y] = val
						nodes[Point{x, y}] = NewNode(x, y, dist)
					} else { // get value from left
						val := fullGrid[x][y-offsetCol]
						if val == 9 {
							val = 1
						} else {
							val++
						}
						fullGrid[x][y] = val
						nodes[Point{x, y}] = NewNode(x, y, dist)
					}
				}
			}
		}
	}

	return fullGrid, nodes
}

func euclideanDistance(grid [][]int, i, j int) float64 {
	m, n := len(grid)-1, len(grid[0])-1
	res := math.Pow(float64(m-i), 2.0) + math.Pow(float64(n-j), 2)
	return math.Sqrt(res)
}

func main() {
	partOne()
	partTwo()
}
