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
	fmt.Println(playGame(2020))
}

func partTwo() {
	fmt.Println(playGame(30000000))
}

func playGame(total int) int {
	startingNums := parseInput()
	turns := make(map[int]int)

	currTurn := 1
	for i := 0; i < len(startingNums)-1; i++ {
		turns[startingNums[i]] = currTurn
		currTurn++
	}

	currTurn++
	lastSpoken := startingNums[len(startingNums)-1]
	lastTurn := currTurn - 1
	currSpoken := 0

	for currTurn <= total {
		if pastTurn, repeat := turns[lastSpoken]; repeat {
			currSpoken = lastTurn - pastTurn
		} else {
			currSpoken = 0
		}
		turns[lastSpoken] = lastTurn
		lastSpoken, lastTurn = currSpoken, currTurn
		currTurn++
	}

	return lastSpoken
}

func parseInput() []int {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal("Failed to open the problem input file")
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	turns := strings.Split(scanner.Text(), ",")
	nums := make([]int, len(turns))
	for i, turn := range turns {
		num, err := strconv.Atoi(turn)
		if err != nil {
			log.Fatal("Failed to convert input to numbers")
		}
		nums[i] = num
	}

	return nums
}

func main() {
	partOne()
	partTwo()
}
