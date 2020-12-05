package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type seat struct {
	row, col int
}

func (s seat) id() int {
	return s.row*8 + s.col
}

func getSeats(inputFilename string) []seat {
	f, _ := os.Open(inputFilename)
	scanner := bufio.NewScanner(f)
	seats := []seat{}

	for scanner.Scan() {
		entry := scanner.Text()
		rowLower, rowUpper := 0, 127
		colLower, colUpper := 0, 7
		row, column := entry[:7], entry[7:]

		for _, letter := range row {
			mid := (rowLower + rowUpper) / 2
			if letter == 'F' {
				rowUpper = mid
			} else {
				rowLower = mid + 1
			}
		}

		for _, letter := range column {
			mid := (colLower + colUpper) / 2
			if letter == 'L' {
				colUpper = mid
			} else {
				colLower = mid + 1
			}
		}

		if colLower != colUpper || rowLower != rowUpper {
			log.Fatal("Err")
		}

		seats = append(seats, seat{rowLower, colLower})
	}

	return seats
}

func partOne(inputFilename string) {
	seats := getSeats(inputFilename)
	maxID := 0
	for _, seat := range seats {
		id := seat.id()
		if id > maxID {
			maxID = id
		}
	}

	fmt.Println(maxID)
}

func partTwo(inpuFilename string) {
	seats := getSeats(inpuFilename)
	candidateIDs := []int{}

	for _, seat := range seats {
		if seat.row != 0 && seat.row != 127 {
			candidateIDs = append(candidateIDs, seat.id())
		}
	}

	sort.Ints(candidateIDs)
	for i, id := range candidateIDs {
		if i > 0 && id != candidateIDs[i-1]+1 {
			fmt.Println(candidateIDs[i-1] + 1)
		}
	}
}

func main() {
	partOne("sample.txt")
	partTwo("input.txt")
}
