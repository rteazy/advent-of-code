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

type Set struct {
	characters map[rune]bool
	word       string
}

func (a Set) intersect(b Set) Set {
	word := []rune{}
	for _, c := range a.word {
		if _, exists := b.characters[c]; exists {
			word = append(word, c)
		}
	}

	return NewSet(string(word))
}

func (a Set) union(b Set) Set {
	chars := []rune{}
	for _, c := range a.word {
		chars = append(chars, c)
	}
	for _, c := range b.word {
		chars = append(chars, c)
	}
	return NewSet(string(chars))
}

func NewSet(word string) Set {
	s := Set{}
	ordered := strings.Split(word, "")
	sort.Strings(ordered)
	s.word = strings.Join(ordered, "")
	s.characters = make(map[rune]bool)
	for _, c := range word {
		s.characters[c] = true
	}
	return s
}

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
		//fmt.Println(val)
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

//wordSets = {
//2: ["ab"],
//3: ["abd"],
//4: ["abef"],
//5: ["bcdef", "acdfg", "abcdf" ],
//6: ["cefabd", "cdfgeb", "cagedb"],
//7: ["abcdefg"]
//}

func buildWordMap(words []string) map[string]int {
	wordMap := make(map[string]int, 0)
	// wordSet maps word length to the set of words
	wordSet := make(map[int][]Set)
	for _, word := range words {
		set := NewSet(word)
		n := len(word)
		if len(wordSet[n]) == 0 {
			wordSet[n] = []Set{set}
		} else {
			wordSet[n] = append(wordSet[n], set)
		}
	}

	// identify 7 segments
	//find posA => for 7 find c <- 1)
	var posA rune
	sevenSet := wordSet[3][0]
	oneSet := wordSet[2][0]
	for _, c := range sevenSet.word {
		if _, exists := oneSet.characters[c]; !exists {
			posA = c
		}
	}
	if posA == 0 {
		log.Fatalf("Could not find character for posA\n")
	}

	//find posB => for 4 find c </- (2 ∩ 3 ∩ 5) U 1
	var posB rune
	fiveLengthIntersection := wordSet[5][0]
	for i := 1; i < len(wordSet[5]); i++ {
		fiveLengthIntersection = fiveLengthIntersection.intersect(wordSet[5][i])
	}
	fiveAndOneUnion := fiveLengthIntersection.union(wordSet[2][0])
	fourSet := wordSet[4][0]
	for _, c := range fourSet.word {
		if _, exists := fiveAndOneUnion.characters[c]; !exists {
			posB = c
			break
		}
	}
	if posB == 0 {
		log.Fatalf("Could not find character for posB\n")
	}

	//find posF => for 5 find c </- {(2 ∩ 3 ∩ 5), b}
	var posF rune

	fiveLengthIntersectionAndB := NewSet(fiveLengthIntersection.union(NewSet(string(posB))).word)
	for _, set := range wordSet[5] {
		missing := 0
		for _, c := range set.word {
			if _, exists := fiveLengthIntersectionAndB.characters[c]; !exists {
				missing++
			}
		}
		if missing == 1 {
			for _, c := range set.word {
				if _, exists := fiveLengthIntersectionAndB.characters[c]; !exists {
					posF = c
				}
			}
		}
	}
	if posF == 0 {
		log.Fatalf("Could not find character for posF\n")
	}

	//find posC => for 1 find c where c != posF
	var posC rune
	for _, c := range oneSet.word {
		if c != posF {
			posC = c
		}
	}
	if posC == 0 {
		log.Fatalf("Could not find character for posC\n")
	}

	//find posE => for 2 find c </- {(2 ∩ 3 ∩ 5), c}
	var posE rune
	fiveLengthIntersectionsAnd4 := fiveLengthIntersection.union(wordSet[4][0])
	for _, set := range wordSet[5] {
		for _, c := range set.word {
			if _, exists := fiveLengthIntersectionsAnd4.characters[c]; !exists {
				posE = c
				break
			}
		}
	}
	if posE == 0 {
		log.Fatalf("Could not find character for posE\n")
	}
	//find posD => for 4 find c </- {c, d, f}
	var posD rune
	bcf := []string{string(posB), string(posC), string(posF)}
	bcfUnion := NewSet(strings.Join(bcf, ""))
	for _, c := range fourSet.word {
		if _, exists := bcfUnion.characters[c]; !exists {
			posD = c
			break
		}
	}
	if posD == 0 {
		log.Fatalf("Could not find character for posD\n")
	}

	//find posG => remaining
	var posG rune
	abcdef := []string{string(posA), string(posB), string(posC), string(posD), string(posE), string(posF)}
	abcdefSet := NewSet(strings.Join(abcdef, ""))
	for _, c := range wordSet[7][0].word {
		if _, exists := abcdefSet.characters[c]; !exists {
			posG = c
		}
	}
	if posG == 0 {
		log.Fatalf("Could not find character for posG\n")
	}
	//fmt.Printf("posA: %s, posB: %s, posC: %s, posD: %s, posE: %s, posF: %s, posG: %s\n", string(posA), string(posB), string(posC), string(posD), string(posE), string(posF), string(posG))

	//wordMap = {"ab": 1, "abef": 4, "abd": 7, .. }
	zeroDigit := NewSet(string([]rune{posA, posB, posC, posE, posF, posG}))
	oneDigit := NewSet(string([]rune{posC, posF}))
	twoDigit := NewSet(string([]rune{posA, posC, posD, posE, posG}))
	threeDigit := NewSet(string([]rune{posA, posC, posD, posF, posG}))
	fourDigit := NewSet(string([]rune{posB, posC, posD, posF}))
	fiveDigit := NewSet(string([]rune{posA, posB, posD, posF, posG}))
	sixDigit := NewSet(string([]rune{posA, posB, posD, posE, posF, posG}))
	sevenDigit := NewSet(string([]rune{posA, posC, posF}))
	eightDigit := NewSet(string([]rune{posA, posB, posC, posD, posE, posF, posG}))
	nineDigit := NewSet(string([]rune{posA, posB, posC, posD, posF, posG}))

	wordMap[zeroDigit.word] = 0
	wordMap[oneDigit.word] = 1
	wordMap[twoDigit.word] = 2
	wordMap[threeDigit.word] = 3
	wordMap[fourDigit.word] = 4
	wordMap[fiveDigit.word] = 5
	wordMap[sixDigit.word] = 6
	wordMap[sevenDigit.word] = 7
	wordMap[eightDigit.word] = 8
	wordMap[nineDigit.word] = 9

	return wordMap
}

func main() {
	partOne()
	partTwo()
}
