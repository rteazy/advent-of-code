package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input.txt", "The relative filepath to the problem input")

type Interval struct {
	start, end int
}

type Ticket struct {
	nums []int
}

func atoiWrap(a string) int {
	val, err := strconv.Atoi(a)
	if err != nil {
		log.Fatalf("Failed to convert %v to int\n", a)
	}
	return val
}

func atoiArray(a string) []int {
	nums := strings.Split(a, ",")
	arr := []int{}
	for _, numString := range nums {
		val, err := strconv.Atoi(numString)
		if err != nil {
			log.Fatalf("Failed to convert %v to int\n", numString)
		}
		arr = append(arr, val)
	}
	return arr
}

func parseInput() (map[string][]Interval, Ticket, []Ticket) {
	fields := make(map[string][]Interval)
	myTicket := Ticket{}
	nearbyTickets := []Ticket{}

	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal("Failed to open the problem input file")
	}

	r, _ := regexp.Compile("^([a-z ]+): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)$")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			scanner.Scan()
			scanner.Scan()
			break
		}
		subMatches := r.FindStringSubmatch(line)
		field := subMatches[1]
		startA, endA, startB, endB := atoiWrap(subMatches[2]), atoiWrap(subMatches[3]), atoiWrap(subMatches[4]), atoiWrap(subMatches[5])
		intervalA, intervalB := Interval{startA, endA}, Interval{startB, endB}
		fields[field] = append(fields[field], intervalA)
		fields[field] = append(fields[field], intervalB)
	}

	myTicketVals := atoiArray(scanner.Text())
	myTicket = Ticket{myTicketVals}

	scanner.Scan()
	scanner.Scan()
	for scanner.Scan() {
		ticketVals := atoiArray(scanner.Text())
		ticket := Ticket{ticketVals}
		nearbyTickets = append(nearbyTickets, ticket)
	}

	return fields, myTicket, nearbyTickets
}

func errorRate(fields map[string][]Interval, tickets []Ticket) {
	invalidVals := []int{}
	for _, ticket := range tickets {
		for _, value := range ticket.nums {
			validValue := false
			for _, intervals := range fields {
				for _, interval := range intervals {
					if !validValue && interval.start <= value && value <= interval.end {
						validValue = true
					}
				}
			}
			if !validValue {
				invalidVals = append(invalidVals, value)
			}
		}
	}

	errorRate := 0
	for _, num := range invalidVals {
		errorRate += num
	}
	fmt.Println(errorRate)
}

func partOne() {
	fields, _, nearbyTickets := parseInput()
	errorRate(fields, nearbyTickets)
}

func pruneInvalidTickets(fields map[string][]Interval, tickets []Ticket) []Ticket {
	validTickets := []Ticket{}
	for _, ticket := range tickets {
		ticketValid := true
		for _, value := range ticket.nums {
			foundInterval := false
			for _, intervals := range fields {
				for _, interval := range intervals {
					if interval.start <= value && value <= interval.end {
						foundInterval = true
					}
				}
			}
			if !foundInterval {
				ticketValid = false
			}
		}
		if ticketValid {
			validTickets = append(validTickets, ticket)
		}
	}

	return validTickets
}

func updateSeatPositions(foundPositions map[string]int, seatPositions []map[string]bool, fields map[string][]Interval, ticket Ticket) {
	for i, val := range ticket.nums {
		for candidate := range seatPositions[i] {
			if _, found := foundPositions[candidate]; found {
				continue
			}
			inRange := false
			for _, interval := range fields[candidate] {
				if !inRange && interval.start <= val && val <= interval.end {
					inRange = true
				}
			}
			if !inRange {
				deleteCandidate(foundPositions, seatPositions, i, candidate)
			}
		}
	}
}

func deleteCandidate(foundPositions map[string]int, seatPositions []map[string]bool, i int, candidate string) {
	delete(seatPositions[i], candidate)
	if len(seatPositions[i]) == 1 {
		remaining := ""
		for k := range seatPositions[i] {
			remaining = k
		}
		foundPositions[remaining] = i

		for j := range seatPositions {
			if _, hasRemaining := seatPositions[j][remaining]; i != j && hasRemaining {
				deleteCandidate(foundPositions, seatPositions, j, remaining)
			}
		}
	}
}

func partTwo() {
	fields, myTicket, nearbyTickets := parseInput()
	validTickets := pruneInvalidTickets(fields, nearbyTickets)

	seatPositions := make([]map[string]bool, len(fields))
	for i := range seatPositions {
		seatPositions[i] = make(map[string]bool)
		for field := range fields {
			seatPositions[i][field] = true
		}
	}

	foundPositions := make(map[string]int)
	for _, ticket := range validTickets {
		updateSeatPositions(foundPositions, seatPositions, fields, ticket)
	}

	// find the positions that start with departure
	indices := []int{}
	for field, i := range foundPositions {
		match, _ := regexp.MatchString("^departure", field)
		if match {
			indices = append(indices, i)
		}
	}

	// return product of myTicket
	product := 1
	for _, pos := range indices {
		product *= myTicket.nums[pos]
	}
	fmt.Println(product)
}

func main() {
	partOne()
	partTwo()
}
