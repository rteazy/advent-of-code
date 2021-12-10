package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
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

	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")[1]
		words := strings.Split(line, " ")
		for _, word := range words {
			n := len(word)
			if n == 2 || n == 4 || n == 3 || n == 7 {
				count++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", count)
}

func partTwo() {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")
		words := strings.Split(line[0], " ")
		wordMap := buildWordMap(words)
		output := strings.Split(line[1], " ")

		digits := make([]string, 0)
		for _, word := range output {
			k := strings.Split(word, "")
			sort.Strings(k)
			num := findNum(strings.Join(k, ""), wordMap)
			digits = append(digits, strconv.Itoa(num))
		}
		val, err := strconv.Atoi(strings.Join(digits, ""))
		if err != nil {
			log.Fatalf("Failed on digits: %v\n", digits)
		}
		total += val
	}

	fmt.Printf("Part 2: %d\n", total)
}

func findNum(word string, words map[string]int) int {
	if _, exists := words[word]; exists {
		return words[word]
	} else {
		log.Fatalf("Could not find mapping for %s\n", word)
	}
	return 0
}

//// map len(word) to numbers
//2: {1}
//3: {7}
//4: {4}
//5: {2, 3, 5}
//6: {0, 6, 9}
//7: {8}

//wordSets = {
//2: ["ab"],
//3: ["abd"],
//4: ["abef"],
//5: ["bcdef", "acdfg", "abcdf" ],
//6: ["cefabd", "cdfgeb", "cagedb"],
//7: ["abcdefg"]
//}
type Set struct {
	characters map[rune]bool
	word       string
}

func NewSet(word string) Set {
	s := Set{}
	s.characters = make(map[rune]bool)
	for _, c := range word {
		s.characters[c] = true
	}
	s.word = word
	return s
}

func buildWordMap(words []string) map[string]int {
	wordMap := make(map[string]int, 0)
	// wordSet maps word length to the set of words
	wordSet := make(map[int][]Set)
	for _, word := range words {
		wordArr := strings.Split(word, "")
		sort.Strings(wordArr)
		set := NewSet(strings.Join(wordArr, ""))
		n := len(wordArr)

		wordSet[n] = append(wordSet[n], set)
	}

	// identify 7 segments
	//find posA => for 7 find c <- 1)
	posA := '_'
	for _, c := range wordSet[3][0].word {
		if _, found := wordSet[1][0].characters[c]; !found {
			posA = c
		}
	}
	fmt.Println(posA)
	//find posB => for 4 find c </- (2 ∩ 3 ∩ 5) U 1
	//find posF => for 5 find c </- {(2 ∩ 3 ∩ 5), b}
	//find posC => for 1 find c where c != posF
	//find posE => for 2 find c </- {(2 ∩ 3 ∩ 5), c}
	//find posD => for 4 find c </- {c, d, f}
	//find posG => remaining

	//wordMap = {"ab": 1, "abef": 4, "abd": 7, .. }
	return wordMap
}

func main() {
	partTwo()
}
