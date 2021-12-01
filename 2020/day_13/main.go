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

func partOne() {
	earliestTime, busIDs := parseInputPartOne()
	curr := earliestTime
	for {
		for _, id := range busIDs {
			if curr%id == 0 {
				diff := curr - earliestTime
				fmt.Println(diff * id)
				return
			}
		}
		curr++
	}
}

func partTwo() {
	busses := parseInputPartTwo()
	busIds := []int{}
	offsets := []int{}
	for i, ch := range strings.Split(busses, ",") {
		if ch != "x" {
			id, err := strconv.Atoi(ch)
			if err != nil {
				log.Fatal("Failed to convert input to busID")
			}
			offsets = append(offsets, id-i)
			busIds = append(busIds, id)
		}
	}

	totalProduct := 1
	for _, id := range busIds {
		totalProduct *= id
	}

	total := 0
	for i := 0; i < len(busIds); i++ {
		n := totalProduct / busIds[i]
		inverse := getInverse(n, busIds[i])
		total += offsets[i] * n * inverse
	}

	fmt.Println(total % totalProduct)
}

func getInverse(a, b int) int {
	start := a % b
	count := 1
	for (start*count)%b != 1 {
		count++
	}
	return count
}

func parseInputPartTwo() string {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal("Failed to open the problem input file")
	}
	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines[1]
}

func parseInputPartOne() (int, []int) {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal("Failed to open the problem input file")
	}
	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	earliestTime, err := strconv.Atoi(lines[0])
	if err != nil {
		log.Fatal("Failed conversion")
	}
	ids := []int{}
	for _, c := range strings.Split(lines[1], ",") {
		if c != "x" {
			busID, err := strconv.Atoi(c)
			if err != nil {
				log.Fatal("Failed conversion")
			}
			ids = append(ids, busID)
		}
	}
	return earliestTime, ids
}

func main() {
	partOne()
	partTwo()
}
