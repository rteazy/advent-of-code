package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var inputFile = flag.String("inputFile", "input.txt", "The relative filepath to the problem input")

func partOne() {
	grid := parseInput()
	prev := -1
	for {
		nextGrid, occupiedSeats := nextState(grid)
		if occupiedSeats == prev {
			fmt.Println(occupiedSeats)
			break
		}

		prev, grid = occupiedSeats, nextGrid
	}
}

func partTwo() {
	grid := parseInput()
	prev := -1
	for {
		nextGrid, occupiedSeats := nextStatePartTwo(grid)
		if occupiedSeats == prev {
			fmt.Println(occupiedSeats)
			break
		}

		prev, grid = occupiedSeats, nextGrid
	}
}

func nextStatePartTwo(grid [][]rune) ([][]rune, int) {
	m, n := len(grid), len(grid[0])
	res := [][]rune{}
	filledSeats := 0

	for i := 0; i < m; i++ {
		newRow := make([]rune, n)
		for j := 0; j < n; j++ {
			directions := [][]int{
				{-1, 0}, {1, 0},
				{0, -1}, {0, 1},
				{-1, -1}, {1, 1},
				{1, -1}, {-1, 1},
			}
			seat := grid[i][j]
			switch seat {
			case 'L':
				occupied := 0
				for _, direction := range directions {
					xOffset, yOffset := direction[0], direction[1]
					x, y := i+xOffset, j+yOffset
					for x >= 0 && x < m && y >= 0 && y < n {
						if grid[x][y] == '#' || grid[x][y] == 'L' {
							if grid[x][y] == '#' {
								occupied++
							}
							break
						}
						x, y = x+xOffset, y+yOffset
					}
				}
				if occupied == 0 {
					newRow[j] = '#'
					filledSeats++
				} else {
					newRow[j] = seat
				}
			case '#':
				occupied := 0
				for _, direction := range directions {
					xOffset, yOffset := direction[0], direction[1]
					x, y := i+xOffset, j+yOffset
					for x >= 0 && x < m && y >= 0 && y < n {
						if grid[x][y] == '#' || grid[x][y] == 'L' {
							if grid[x][y] == '#' {
								occupied++
							}
							break
						}
						x, y = x+xOffset, y+yOffset
					}
				}
				if occupied >= 5 {
					newRow[j] = 'L'
				} else {
					newRow[j] = seat
					filledSeats++
				}
			case '.':
				newRow[j] = seat
			default:
				log.Fatal("Failed")
			}
		}
		res = append(res, newRow)
	}

	return res, filledSeats
}

func nextState(grid [][]rune) ([][]rune, int) {
	m, n := len(grid), len(grid[0])
	res := [][]rune{}
	filledSeats := 0

	for i := 0; i < m; i++ {
		newRow := make([]rune, n)
		for j := 0; j < n; j++ {
			neighbors := [][]int{
				{i - 1, j}, {i + 1, j},
				{i, j - 1}, {i, j + 1},
				{i - 1, j - 1}, {i + 1, j + 1},
				{i + 1, j - 1}, {i - 1, j + 1},
			}
			seat := grid[i][j]
			switch seat {
			case 'L':
				occupied := 0
				for _, neighbor := range neighbors {
					x, y := neighbor[0], neighbor[1]
					if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == '#' {
						occupied++
					}
				}
				if occupied == 0 {
					newRow[j] = '#'
					filledSeats++
				} else {
					newRow[j] = seat
				}
			case '#':
				occupied := 0
				for _, neighbor := range neighbors {
					x, y := neighbor[0], neighbor[1]
					if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == '#' {
						occupied++
					}
				}
				if occupied >= 4 {
					newRow[j] = 'L'
				} else {
					newRow[j] = seat
					filledSeats++
				}
			case '.':
				newRow[j] = seat
			default:
				log.Fatal("Failed")
			}
		}
		res = append(res, newRow)
	}

	return res, filledSeats
}

func printGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
	fmt.Println("-------------------")
}

func parseInput() [][]rune {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal("Failed to open the problem input")
	}

	grid := [][]rune{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		grid = append(grid, []rune(row))
	}
	return grid
}

func main() {
	//partOne()
	partTwo()
}
