package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var inputFile = flag.String("inputFile", "input.txt", "The relative filepath to the problem input")

func partOne() {
	actions := parseInput()
	x, y := 0, 0
	directions := []rune{'N', 'E', 'S', 'W'}
	dirIndex := 1
	for _, act := range actions {
		switch act.direction {
		case 'N':
			y += act.value
		case 'S':
			y -= act.value
		case 'W':
			x -= act.value
		case 'E':
			x += act.value
		case 'L':
			offset := (act.value / 90) % len(directions)
			dirIndex -= offset
			if dirIndex < 0 {
				dirIndex = len(directions) + dirIndex
			}
		case 'R':
			dirIndex = (dirIndex + (act.value / 90)) % len(directions)
		case 'F':
			currDirection := directions[dirIndex]
			switch currDirection {
			case 'E':
				x += act.value
			case 'W':
				x -= act.value
			case 'N':
				y += act.value
			case 'S':
				y -= act.value
			}
		}
	}
	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func partTwo() {
	actions := parseInput()
	waypointX, waypointY := 10.0, 1.0
	positionX, positionY := 0.0, 0.0

	for _, act := range actions {
		val := float64(act.value)
		switch act.direction {
		case 'N':
			waypointY += val
		case 'S':
			waypointY -= val
		case 'W':
			waypointX -= val
		case 'E':
			waypointX += val
		case 'L', 'R':
			rotationRadians := val * (math.Pi / 180.0)
			if act.direction == 'R' {
				rotationRadians *= -1
			}
			newWayPointX := waypointX*math.Cos(rotationRadians) - waypointY*math.Sin(rotationRadians)
			newWayPointY := waypointY*math.Cos(rotationRadians) + waypointX*math.Sin(rotationRadians)
			waypointX, waypointY = newWayPointX, newWayPointY
		case 'F':
			positionX += waypointX * val
			positionY += waypointY * val
		}
	}

	fmt.Println(math.Abs(positionX) + math.Abs(positionY))
}

type action struct {
	direction rune
	value     int
}

func parseInput() []action {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal("Failed to open the problem input file")
	}

	actions := []action{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		dir, val := line[0], line[1:]
		value, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal("Failed to convert the input")
		}
		actions = append(actions, action{rune(dir), value})
	}
	return actions
}

func main() {
	partOne()
	partTwo()
}
