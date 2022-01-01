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

type Point struct {
	x, y int
}

func partOne() {
	points, _, _ := run(true)
	fmt.Printf("Part 1: %d\n", len(points))
}

func partTwo() {
	points, endX, endY := run(false)
	fmt.Println("Part 2: ")
	printPoints(points, endX, endY)
}

func run(exitAfterFirstFold bool) (map[Point]bool, int, int) {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	points := make(map[Point]bool)
	endX, endY := -1, -1
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		tokens := strings.Split(line, ",")
		x, err := strconv.Atoi(tokens[0])
		if err != nil {
			log.Fatalf("Failed conversino for %s\n", tokens[0])
		}
		y, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatalf("Failed conversino for %s\n", tokens[1])
		}
		if x > endX {
			endX = x
		}
		if y > endY {
			endY = y
		}
		point := Point{x, y}
		points[point] = true
	}
	for scanner.Scan() {
		line := scanner.Text()
		text := strings.Split(line, " ")[2]
		instructions := strings.Split(text, "=")
		axis := instructions[0]
		val, err := strconv.Atoi(instructions[1])
		if err != nil {
			log.Fatalf("Failed for %s\n", instructions[1])
		}
		if axis == "x" {
			endX = foldXAxis(val, points, endY)
		} else if axis == "y" {
			endY = foldYAxis(val, points, endX)
		} else {
			log.Fatalf("Unexpected: %s\n", axis)
		}
		if exitAfterFirstFold {
			break
		}
	}

	return points, endX, endY
}

func foldXAxis(verticalX int, points map[Point]bool, endY int) int {
	i := 0
	endX := verticalX * 2
	for x := endX; x > verticalX; x-- {
		for y := 0; y <= endY; y++ {
			point := Point{x, y}
			if _, exists := points[point]; exists {
				delete(points, point)
				newPoint := Point{i, y}
				points[newPoint] = true
			}
		}
		i++
	}

	return verticalX - 1
}

func foldYAxis(horizontalY int, points map[Point]bool, endX int) int {
	i := 0
	endY := horizontalY * 2
	for y := endY; y > horizontalY; y-- {
		for x := 0; x <= endX; x++ {
			point := Point{x, y}
			if _, exists := points[point]; exists {
				delete(points, point)
				newPoint := Point{x, i}
				points[newPoint] = true
			}
		}
		i++
	}
	return horizontalY - 1
}

func printPoints(points map[Point]bool, endX, endY int) {
	for y := 0; y <= endY; y++ {
		for x := 0; x <= endX; x++ {
			p := Point{x, y}
			if _, exists := points[p]; exists {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	partOne()
	partTwo()
}
