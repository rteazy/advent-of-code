package main

import (
	"bufio"
	"fmt"
	"os"
)

func partTwo(inputFilename string) {
	res := 0
	groupAnswers, peoplePerGroup := getGroupAnswers(inputFilename)
	for i, answers := range groupAnswers {
		for _, freq := range answers {
			if freq == peoplePerGroup[i] {
				res++
			}
		}
	}
	fmt.Println(res)
}

func partOne(inputFilename string) {
	groupsAnswers, _ := getGroupAnswers(inputFilename)
	counts := 0
	for _, group := range groupsAnswers {
		counts += len(group)
	}

	fmt.Println(counts)
}

func getGroupAnswers(inputFilename string) ([]map[rune]int, []int) {
	f, _ := os.Open(inputFilename)
	scanner := bufio.NewScanner(f)
	groups := []map[rune]int{}
	groups = append(groups, make(map[rune]int))
	peoplePerGroup := []int{}
	numPeople := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			peoplePerGroup = append(peoplePerGroup, numPeople)
			numPeople = 0
			counts := make(map[rune]int)
			groups = append(groups, counts)
		} else {
			numPeople++
			for _, c := range line {
				lastGroup := len(groups) - 1
				if _, exists := groups[lastGroup][c]; !exists {
					groups[lastGroup][c] = 1
				} else {
					groups[lastGroup][c]++
				}
			}
		}
	}

	peoplePerGroup = append(peoplePerGroup, numPeople)

	return groups, peoplePerGroup
}

func main() {
	partOne("sample.txt")
	partOne("input.txt")
	partTwo("sample.txt")
	partTwo("input.txt")
}
