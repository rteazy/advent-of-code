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

type setMemoryFunc func(string, map[uint64]uint64, uint64, uint64)

func partTwo() {
	countTotalValuesInMemory(setMemoryPart2)
}

func partOne() {
	countTotalValuesInMemory(setMemoryPart1)
}

func countTotalValuesInMemory(f setMemoryFunc) {
	addresses := make(map[uint64]uint64)
	currMask := ""
	lines := parseInput()

	for _, line := range lines {
		exp := strings.Split(line, " = ")
		lhs, rhs := exp[0], exp[1]
		if lhs == "mask" {
			currMask = rhs
		} else {
			addStr := lhs[4 : len(lhs)-1]
			address, err := strconv.ParseUint(addStr, 10, 64)
			if err != nil {
				log.Fatalf("Failed to convert %v to decimal", addStr)
			}

			value, err := strconv.ParseUint(rhs, 10, 64)
			if err != nil {
				log.Fatalf("Failed to convert %v to decimal", rhs)
			}

			f(currMask, addresses, address, value)
		}
	}

	var total uint64
	for _, val := range addresses {
		total += val
	}
	fmt.Println(total)
}

func setMemoryPart2(bitmask string, addresses map[uint64]uint64, currAddress, value uint64) {
	// given the bitmask, perform OR
	addressBinary := strconv.FormatUint(currAddress, 2)
	for len(addressBinary) != 36 {
		addressBinary = "0" + addressBinary
	}

	result := make([]rune, len(addressBinary))
	for i := 0; i < len(addressBinary); i++ {
		addressBit := bitmask[i]
		if addressBit == '0' {
			result[i] = rune(addressBinary[i])
		} else {
			result[i] = rune(addressBit)
		}
	}

	// generate the address combinations given the result
	combinations := generateAddressCombinations(string(result))

	// for each address in the combinations, set the value for the address
	for _, address := range combinations {
		addresses[address] = value
	}
}

func generateAddressCombinations(result string) []uint64 {
	combinations := make(map[uint64]bool)
	helper(result, 0, "", combinations)
	addresses := []uint64{}
	for address := range combinations {
		addresses = append(addresses, address)
	}
	return addresses
}

func helper(result string, currIndex int, addressString string, combinations map[uint64]bool) {
	if currIndex == len(result) {
		address, err := strconv.ParseUint(addressString, 2, 64)
		if err != nil {
			log.Fatalf("Failed to get binary number from mask: %s\n", addressString)
		}
		combinations[address] = true
		return
	}
	if result[currIndex] == 'X' {
		helper(result, currIndex+1, addressString+"0", combinations)
		helper(result, currIndex+1, addressString+"1", combinations)
	} else {
		helper(result, currIndex+1, addressString+string(result[currIndex]), combinations)
	}
}

func setMemoryPart1(bitmask string, addresses map[uint64]uint64, address, value uint64) {
	bitMaskRep := []rune{}
	bitPositions := []int{}

	for i, c := range bitmask {
		if c == '0' || c == '1' {
			bitPositions = append(bitPositions, len(bitmask)-1-i)
			bitMaskRep = append(bitMaskRep, c)
		} else {
			bitMaskRep = append(bitMaskRep, '0')
		}
	}

	mask, err := strconv.ParseUint(string(bitMaskRep), 2, 64)
	if err != nil {
		log.Fatalf("Failed to get binary number from mask: %s\n", string(bitMaskRep))
	}

	for _, pos := range bitPositions {
		var val uint64
		val = 1
		val <<= pos
		value &= ^val
	}

	result := value | mask
	addresses[address] = result
}

func parseInput() []string {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal("Failed to open the problem input file")
	}

	scanner := bufio.NewScanner(f)
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func main() {
	partOne()
	partTwo()
}
