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

var inputFile = flag.String("inputFile", "input.txt", "relative filepath to the problem input")

type instruction struct {
	op  string
	val int
}

func partOne() {
	instructions := parseInstructions()
	_, acc := runInstructions(instructions)
	fmt.Println(acc)
}

func partTwo() {
	instructions := parseInstructions()
	instructionsCombinations := [][]instruction{instructions}

	for i, in := range instructions {
		if in.op == "jmp" || in.op == "op" {
			cpy := make([]instruction, len(instructions))
			copy(cpy, instructions)
			if in.op == "jmp" {
				cpy[i] = instruction{"nop", in.val}
			} else {
				cpy[i] = instruction{"jmp", in.val}
			}
			instructionsCombinations = append(instructionsCombinations, cpy)
		}

		for _, comb := range instructionsCombinations {
			terminates, acc := runInstructions(comb)
			if terminates {
				fmt.Println(acc)
				return
			}
		}
	}
}

func runInstructions(instructions []instruction) (bool, int) {
	accumulator := 0
	curr := 0
	seen := make(map[int]bool)
	terminated := false

	for {
		if _, cycle := seen[curr]; cycle {
			break
		}
		if curr == len(instructions) {
			terminated = true
			break
		}

		seen[curr] = true
		switch instructions[curr].op {
		case "jmp":
			curr += instructions[curr].val
		case "acc":
			accumulator += instructions[curr].val
			curr++
		default:
			curr++
		}
	}

	return terminated, accumulator
}

func parseInstructions() []instruction {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal("Failed to open input file")
	}
	scanner := bufio.NewScanner(f)
	instructions := []instruction{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		op := line[0]
		value, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatalf("Failed parsing operation")
		}
		instructions = append(instructions, instruction{op, value})
	}

	return instructions
}

func main() {
	//partOne()
	partTwo()
}
