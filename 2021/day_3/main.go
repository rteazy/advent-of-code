package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

var inputFile = flag.String("inputFile", "input.txt", "The relative filepath to the problem input")

func partOne() {
	data := parseInput()
	numBits := len(data[0])
	counter := make([]int, numBits)
	for _, bits := range data {
		for i, val := range bits {
			counter[i] += val
		}
	}

	ep, gam := make([]int, numBits), make([]int, numBits)
	for i, count := range counter {
		if count > len(data)/2 {
			ep[i] = 1
			gam[i] = 0
		} else {
			ep[i] = 0
			gam[i] = 1
		}
	}

	fmt.Printf("Part One: %d\n", convertToDec(ep)*convertToDec(gam))
}

func partTwo() {
	data := parseInput()
	a := findNumber(data, true)
	b := findNumber(data, false)
	fmt.Printf("Part Two: %d\n", a*b)
}

func findNumber(data [][]int, mostCommon bool) int {
	candidates := data
	bitLength := len(data[0])
	i := 0
	for len(candidates) != 1 && i < bitLength {
		countOnes := 0
		ones, zeros := make([][]int, 0), make([][]int, 0)
		for _, bits := range candidates {
			countOnes += bits[i]
			if bits[i] == 1 {
				ones = append(ones, bits)
			} else {
				zeros = append(zeros, bits)
			}
		}

		countZeros := len(candidates) - countOnes
		if mostCommon {
			if countOnes < countZeros {
				candidates = zeros
			} else {
				candidates = ones
			}
		} else {
			if countOnes < countZeros {
				candidates = ones
			} else {
				candidates = zeros
			}
		}

		i++
	}

	if len(candidates) != 1 {
		log.Fatalf("Failed to find single candidate")
	}

	return convertToDec(candidates[0])
}

func parseInput() [][]int {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	scanner := bufio.NewScanner(f)
	inputData := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		data := make([]int, len(line))
		for i, c := range line {
			data[i] = int(c - '0')
		}
		inputData = append(inputData, data)
	}

	return inputData
}

func convertToDec(bits []int) int {
	total := 0
	idx := len(bits) - 1
	for i := idx; i >= 0; i-- {
		if bits[i] == 1 {
			total += int(math.Pow(2, float64(idx-i)))
		}
	}
	return total
}

func main() {
	partOne()
	partTwo()
}
