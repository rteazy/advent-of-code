package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partTwo(inputFilename string) {
	records := parseInput(inputFilename)
	validPasswords := []string{}

	for _, record := range records {
		iStr, jStr, letter, password := record[0], record[1], []rune(record[2])[0], record[3]

		i, err := strconv.ParseInt(iStr, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		j, err := strconv.ParseInt(jStr, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		a, b := rune(password[i-1]), rune(password[j-1])

		if a == letter && b == letter {
			continue
		} else if a == letter || b == letter {
			validPasswords = append(validPasswords, password)
		}
	}

	fmt.Println(len(validPasswords))
}

func partOne(inputFilename string) {
	records := parseInput(inputFilename)
	res := []string{}

	for _, record := range records {
		minFreqStr, maxFrexStr, letter, password := record[0], record[1], []rune(record[2])[0], record[3]
		counter := make(map[rune]int64)
		for _, r := range password {
			counter[r]++
		}

		minFreq, err := strconv.ParseInt(minFreqStr, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		maxFreq, err := strconv.ParseInt(maxFrexStr, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		if _, exists := counter[letter]; exists && minFreq <= counter[letter] && counter[letter] <= maxFreq {
			res = append(res, password)
		}
	}

	fmt.Println(len(res))
}

func parseInput(fileName string) [][]string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data := [][]string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		record := strings.Split(scanner.Text(), " ")
		frequency, letter, password := record[0], record[1], record[2]

		freq := strings.Split(frequency, "-")
		minFreq, maxFreq := freq[0], freq[1]

		data = append(data, []string{minFreq, maxFreq, letter, password})
	}

	return data
}

func main() {
	partOne("sample.txt")
	partOne("input.txt")
	partTwo("sample.txt")
	partTwo("input.txt")
}
