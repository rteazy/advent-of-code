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
	i, j int
}

type Board struct {
	nums   [][]int
	marked map[Point]bool
	won    bool
}

func NewBoard() *Board {
	nums := make([][]int, 0)
	marked := make(map[Point]bool, 0)
	return &Board{nums, marked, false}
}

func (b *Board) check(num int) bool {
	m, n := len(b.nums), len(b.nums[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if b.nums[i][j] == num {
				b.marked[Point{i, j}] = true
				verticalCount := 0
				horizontalCount := 0
				for x := 0; x < m; x++ {
					p := Point{x, j}
					if _, exists := b.marked[p]; exists {
						verticalCount++
					}
				}
				for y := 0; y < n; y++ {
					p := Point{i, y}
					if _, exists := b.marked[p]; exists {
						horizontalCount++
					}
				}
				if verticalCount == m || horizontalCount == n {
					b.won = true
					return true
				}
			}

		}
	}
	return false
}

func parseInput() ([]int, []*Board) {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	drawnStr := strings.Split(scanner.Text(), ",")
	drawn := make([]int, len(drawnStr))
	for i, s := range drawnStr {
		c, _ := strconv.Atoi(s)
		drawn[i] = c
	}
	scanner.Scan()

	boards := make([]*Board, 0)
	board := NewBoard()
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			boards = append(boards, board)
			board = NewBoard()
			continue
		}
		line := scanner.Text()
		nums := make([]int, 0)
		num := ""
		for _, c := range line {
			if c == ' ' {
				if num != "" {
					v, err := strconv.Atoi(num)
					if err != nil {
						log.Fatalf("Failed to convert %s\n", num)
					}
					nums = append(nums, v)
					num = ""
				}
				continue
			} else {
				num += string(c)
			}
		}
		v, err := strconv.Atoi(num)
		if err != nil {
			log.Fatalf("Failed to convert %s\n", num)
		}
		nums = append(nums, v)
		board.nums = append(board.nums, nums)
	}
	boards = append(boards, board)
	return drawn, boards
}

func partOne() {
	nums, boards := parseInput()
	for _, num := range nums {
		for _, board := range boards {
			if bingo := board.check(num); bingo {
				total := 0
				for i := 0; i < len(board.nums); i++ {
					for j := 0; j < len(board.nums[0]); j++ {
						p := Point{i, j}
						if _, marked := board.marked[p]; !marked {
							total += board.nums[i][j]
						}
					}
				}
				fmt.Printf("Part 1: %d\n", num*total)
				return
			}
		}
	}
}

func partTwo() {
	nums, boards := parseInput()
	lastScore := 0
	for _, num := range nums {
		for _, board := range boards {
			if !board.won {
				if bingo := board.check(num); bingo {
					//fmt.Printf("Bingo: %d\n", num)
					total := 0
					for i := 0; i < len(board.nums); i++ {
						for j := 0; j < len(board.nums[0]); j++ {
							p := Point{i, j}
							if _, marked := board.marked[p]; !marked {
								total += board.nums[i][j]
							}
						}
					}
					lastScore = total * num
				}
			}
		}
	}
	fmt.Printf("Part 2: %d\n", lastScore)
}

func main() {
	partOne()
	partTwo()
}
