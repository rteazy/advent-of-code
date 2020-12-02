package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	sampleData := readData("sample.txt")
	inputData := readData("input.txt")

	partOne(sampleData)
	partOne(inputData)

	partTwo(sampleData)
	partTwo(inputData)
}

func twoSum(nums []int, c int) (int, int, error) {
	seen := make(map[int]bool)
	for _, a := range nums {
		b := c - a
		if _, ok := seen[b]; ok {
			return a, b, nil
		}

		seen[a] = true
	}

	return -1, -1, errors.New("No pair found that sums to target")
}

func threeSum(nums []int, target int) (int, int, int, error) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				a, b, c := nums[i], nums[j], nums[k]
				if a+b+c == target {
					return a, b, c, nil
				}
			}
		}
	}
	return -1, -1, -1, errors.New("No pair found that sums to target")
}

func partOne(data []int) {
	a, b, err := twoSum(data, 2020)
	check(err)
	fmt.Printf("%d\n", a*b)
}

func partTwo(data []int) {
	a, b, c, err := threeSum(data, 2020)
	check(err)
	fmt.Printf("%d\n", a*b*c)
}

func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readData(inputFile string) []int {
	f, err := os.Open(inputFile)
	check(err)
	data, err := readInts(bufio.NewReader(f))
	check(err)
	return data
}
