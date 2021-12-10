package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

var inputFile = flag.String("inputFile", "input.txt", "The relative filepath to the problem input")

func partOne() {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	scanner := bufio.NewScanner(f)
	invalidScore := make(map[rune]int)
	invalidScore[')'] = 3
	invalidScore[']'] = 57
	invalidScore['}'] = 1197
	invalidScore['>'] = 25137

	complement := make(map[rune]rune)
	complement['('] = ')'
	complement['['] = ']'
	complement['{'] = '}'
	complement['<'] = '>'
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		stack := make([]rune, 0)
		for _, c := range line {
			if len(stack) == 0 {
				stack = append(stack, c)
			} else {
				if _, open := complement[c]; open {
					stack = append(stack, c)
				} else {
					if complement[stack[len(stack)-1]] != c {
						total += invalidScore[c]
						break
					} else {
						stack = stack[:len(stack)-1]
					}
				}
			}

		}
	}
	fmt.Printf("Part 1: %d\n", total)
}

func partTwo() {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	scanner := bufio.NewScanner(f)
	complement := make(map[rune]rune)
	complement['('] = ')'
	complement['['] = ']'
	complement['{'] = '}'
	complement['<'] = '>'
	result := make([]int, 0)

	for scanner.Scan() {
		invalidLine := false
		line := scanner.Text()
		stack := make([]rune, 0)
		for _, c := range line {
			if len(stack) == 0 {
				stack = append(stack, c)
			} else {
				if _, open := complement[c]; open {
					stack = append(stack, c)
				} else {
					n := len(stack) - 1
					if complement[stack[n]] != c {
						// invalid line, discard
						invalidLine = true
						break
					} else {
						stack = stack[:n]
					}
				}
			}
		}
		if invalidLine { //move to the next line
			continue
		}
		//incomplete
		if len(stack) > 0 {
			closing := make([]rune, 0)
			for len(stack) > 0 {
				lastIndex := len(stack) - 1
				closing = append(closing, complement[stack[lastIndex]])
				stack = stack[:lastIndex]
			}
			result = append(result, score(closing))
		}
	}

	sort.Ints(result)
	fmt.Printf("Part 2: %d\n", result[len(result)/2])
}

func score(closing []rune) int {
	total := 0
	scoreMap := make(map[rune]int)
	scoreMap[')'] = 1
	scoreMap[']'] = 2
	scoreMap['}'] = 3
	scoreMap['>'] = 4
	for _, c := range closing {
		total = total*5 + scoreMap[c]
	}
	return total
}

func main() {
	partOne()
	partTwo()
}
