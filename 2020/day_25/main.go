package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var inputFile = flag.String("inputFile", "input.txt", "The relative filepath to the problem input")
var debug = flag.Bool("debug", false, "Debug information")

func partOne() {
	parseInput()
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal("Failed to open the problem input file")
	}

	scanner := bufio.NewScanner(f)
	publicKeys := []int{}
	for scanner.Scan() {
		publicKey, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Failed to parse the public key")
		}
		publicKeys = append(publicKeys, publicKey)
	}
	// Given public key, find the loop size
	loopSizeA := findLoopSize(publicKeys[0])
	loopSizeB := findLoopSize(publicKeys[1])

	// Transform to get the public encryption
	fmt.Println(transform(publicKeys[0], loopSizeB))
	fmt.Println(transform(publicKeys[1], loopSizeA))
}

func findLoopSize(publicKey int) int {
	subjectNumber := 7
	val := 1
	loopSize := 0
	for val != publicKey {
		val *= subjectNumber
		val %= 20201227
		loopSize++
	}
	return loopSize
}

func transform(subjectNumber, loopSize int) int {
	if *debug {
		fmt.Printf("Loop size: %d, Subject number: %d\n", loopSize, subjectNumber)
	}
	val := 1
	for i := 0; i < loopSize; i++ {
		val *= subjectNumber
		val %= 20201227
	}
	return val
}

func parseInput() {
}

func main() {
	partOne()
}
