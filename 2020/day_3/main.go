package main

import (
	"bufio"
	"fmt"
	"os"
)

func partOne(grid [][]rune) {
	fmt.Println(numTrees(grid, 3, 1))
}

func partTwo(grid [][]rune) {
	res := 1
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	for _, slope := range slopes {
		right, down := slope[0], slope[1]
		res *= numTrees(grid, right, down)
	}

	fmt.Println(res)
}

func numTrees(grid [][]rune, right, down int) int {
	const tree = '#'

	width := len(grid[0])
	count := 0
	j := 0

	for i := 0; i < len(grid); i += down {
		row := grid[i]
		if row[j] == tree {
			count++
		}
		j = (j + right) % width
	}
	return count
}

func createGrid(inputFilename string) [][]rune {
	f, _ := os.Open(inputFilename)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	grid := [][]rune{}
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	return grid
}

func main() {
	testInputGrid := createGrid("sample.txt")
	inputGrid := createGrid("input.txt")

	partOne(testInputGrid)
	partOne(inputGrid)
	partTwo(testInputGrid)
	partTwo(inputGrid)
}
