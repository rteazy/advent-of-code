package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var inputFile = flag.String("inputFile", "input.txt", "relative filepath to problem input")
var size = flag.Int("size", 25, "The length of the preamble")

func partOne() {
	fmt.Println(findInvalidNumber(*size, parseInput()))
}

func partTwo() {
	data := parseInput()
	invalidNumber := findInvalidNumber(*size, data)

	i := 0
	total := 0

	for j := i; j < len(data); j++ {
		if total == invalidNumber {
			minVal, maxVal := data[i], data[i]
			for k := i; k < j; k++ {
				num := data[k]
				if num < minVal {
					minVal = num
				}
				if num > maxVal {
					maxVal = num
				}
			}
			fmt.Println(minVal + maxVal)
			return
		}

		total += data[j]
		for total > invalidNumber && i < j {
			total -= data[i]
			i++
		}
	}
}

func findInvalidNumber(preambleSize int, data []int) int {
	seen := make(map[int]bool)
	i := 0
	for i = 0; i < preambleSize; i++ {
		seen[data[i]] = true
	}

	for i < len(data) {
		target, found := data[i], false
		for num := range seen {
			if _, pairFound := seen[target-num]; pairFound {
				found = true
				break
			}
		}
		if !found {
			return target
		}

		delete(seen, data[i-preambleSize])
		i++
		seen[target] = true
	}

	return -1
}

func parseInput() []int {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal("Failed to open the input file")
	}

	scanner := bufio.NewScanner(f)
	data := []int{}
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Failed to convert")
		}
		data = append(data, num)
	}

	return data
}

func main() {
	partOne()
	partTwo()
}
