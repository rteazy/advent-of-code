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
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	scanner := bufio.NewScanner(f)
	counts := make(map[Coordinate]int)
	for scanner.Scan() {
		line := scanner.Text()
		pairs := strings.Split(line, " -> ")
		startStr := strings.Split(pairs[0], ",")
		x1, _ := strconv.Atoi(startStr[0])
		y1, _ := strconv.Atoi(startStr[1])
		endStr := strings.Split(pairs[1], ",")
		x2, _ := strconv.Atoi(endStr[0])
		y2, _ := strconv.Atoi(endStr[1])
		//fmt.Printf("(%d, %d) -> (%d, %d}\n", x1, y1, x2, y2)
		if x1 == x2 {
			x := x1
			startY, endY := 0, 0
			if y1 < y2 {
				startY, endY = y1, y2
			} else {
				startY, endY = y2, y1
			}
			for i := startY; i <= endY; i++ {
				addToCounts(counts, Coordinate{x, i})
			}
		} else if y1 == y2 {
			y := y1
			startX, endX := 0, 0
			if x1 < x2 {
				startX, endX = x1, x2
			} else {
				startX, endX = x2, x1
			}
			for i := startX; i <= endX; i++ {
				addToCounts(counts, Coordinate{i, y})
			}
		}
	}
	res := 0
	for _, v := range counts {
		if v >= 2 {
			res++
		}
	}
	fmt.Printf("Part 1: %d\n", res)
}

//a - b
//- i -
//c - d
//
//i = (x1, y1)
//j = (x2, y2)
//
//a = x1 - 1, y1 - 1  if x2 < x1 && y2 < y1
//b = x1 + 1, y1 - 1  if x2 > x1 && y2 < y1
//c = x1 - 1, y1 + 1  if x2 < x1 && y2 > y1
//d = x1 + 1, y1 + 1  if x2 > x1 && y2 > y1

func partTwo() {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	scanner := bufio.NewScanner(f)
	counts := make(map[Coordinate]int)
	for scanner.Scan() {
		line := scanner.Text()
		pairs := strings.Split(line, " -> ")
		startStr := strings.Split(pairs[0], ",")
		x1, _ := strconv.Atoi(startStr[0])
		y1, _ := strconv.Atoi(startStr[1])
		endStr := strings.Split(pairs[1], ",")
		x2, _ := strconv.Atoi(endStr[0])
		y2, _ := strconv.Atoi(endStr[1])
		if x1 == x2 {
			x := x1
			startY, endY := 0, 0
			if y1 < y2 {
				startY, endY = y1, y2
			} else {
				startY, endY = y2, y1
			}
			for i := startY; i <= endY; i++ {
				addToCounts(counts, Coordinate{x, i})
			}
		} else if y1 == y2 {
			y := y1
			startX, endX := 0, 0
			if x1 < x2 {
				startX, endX = x1, x2
			} else {
				startX, endX = x2, x1
			}
			for i := startX; i <= endX; i++ {
				addToCounts(counts, Coordinate{i, y})
			}
		} else {
			dirX := 0
			dirY := 0
			if x2 < x1 && y2 < y1 {
				dirX, dirY = -1, -1
			} else if x2 > x1 && y2 < y1 {
				dirX, dirY = 1, -1
			} else if x2 < x1 && y2 > y1 {
				dirX, dirY = -1, 1
			} else if x2 > x1 && y2 > y1 {
				dirX, dirY = 1, 1
			} else {
				log.Fatalf("Failed on (%d,%d)->(%d,%d)\n", x1, y1, x2, y2)
			}
			currX, currY := x1, y1
			for currX != x2 && currY != y2 {
				addToCounts(counts, Coordinate{currX, currY})
				currX += dirX
				currY += dirY
			}
			addToCounts(counts, Coordinate{currX, currY})
		}
	}
	res := 0
	for _, v := range counts {
		if v >= 2 {
			res++
		}
	}
	fmt.Printf("Part 2: %d\n", res)
}

func addToCounts(counts map[Coordinate]int, coord Coordinate) {
	if _, exists := counts[coord]; exists {
		counts[coord]++
	} else {
		counts[coord] = 1
	}
}

func main() {
	partOne()
	partTwo()
}
