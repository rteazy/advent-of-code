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
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	horizontal := 0
	depth := 0
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		lineStr := strings.Split(line, " ")
		dir := lineStr[0]
		val, _ := strconv.Atoi(lineStr[1])
		switch dir {
		case "forward":
			horizontal += val
		case "down":
			depth += val
		case "up":
			depth -= val
		}
	}
	fmt.Println(depth * horizontal)
}
func partTwo() {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	horizontal := 0
	depth := 0
	aim := 0
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		lineStr := strings.Split(line, " ")
		dir := lineStr[0]
		val, _ := strconv.Atoi(lineStr[1])
		switch dir {
		case "forward":
			horizontal += val
			depth = depth + (aim * val)
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}
	fmt.Println(depth * horizontal)
}

func main() {
	partOne()
	partTwo()
}
