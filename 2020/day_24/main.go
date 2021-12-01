package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var inputFile = flag.String("inputFile", "input.txt", "The relative filepath to the problem input")
var numDays = flag.Int("numDays", 100, "The number of days passed")

func partOne() {
	colors := getInitialColors()

	numBlack := 0
	for _, color := range colors {
		numBlack += color
	}

	fmt.Println(numBlack)
}

func getInitialColors() map[Coordinate]int {
	flag.Parse()
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal("Failed to open the problem input file")
	}

	tiles := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		tiles = append(tiles, scanner.Text())
	}

	colors := make(map[Coordinate]int)
	for _, tile := range tiles {
		x, y, z := 0, 0, 0
		for i := 0; i < len(tile); i++ {
			switch tile[i] {
			case 'w':
				x--
				y++
			case 'e':
				x++
				y--
			case 'n':
				i++
				if tile[i] == 'e' {
					x++
					z--
				} else {
					y++
					z--
				}
			case 's':
				i++
				if tile[i] == 'e' {
					y--
					z++
				} else {
					x--
					z++
				}
			default:
				log.Fatalf("Unexpected: %v", tile[i])
			}
		}

		coord := Coordinate{x, y, z}
		if _, exists := colors[coord]; exists {
			colors[coord] ^= 1
		} else {
			colors[coord] = 1 //mark black
		}
	}

	offsets := [][]int{{1, -1, 0}, {-1, 1, 0}, {0, -1, 1}, {-1, 0, 1}, {1, 0, -1}, {0, 1, -1}}
	for tile := range colors {
		for _, offset := range offsets {
			x, y, z := tile.x+offset[0], tile.y+offset[1], tile.z+offset[2]
			neighbor := Coordinate{x, y, z}
			if _, exists := colors[neighbor]; !exists {
				colors[neighbor] = 0 // if by default
			}
		}
	}

	return colors
}

func countAdjacentBlackTiles(tile Coordinate, tiles map[Coordinate]int, nextState map[Coordinate]int) int {
	// e w se sw ne nw
	offsets := [][]int{{1, -1, 0}, {-1, 1, 0}, {0, -1, 1}, {-1, 0, 1}, {1, 0, -1}, {0, 1, -1}}
	blackAdjCount := 0
	for _, offset := range offsets {
		x, y, z := tile.x+offset[0], tile.y+offset[1], tile.z+offset[2]
		neighbor := Coordinate{x, y, z}
		if _, exists := tiles[neighbor]; exists {
			blackAdjCount += tiles[neighbor]
		}
		// add neighbors to the next state
		if _, exists := nextState[neighbor]; !exists {
			nextState[neighbor] = 0
		}
	}
	return blackAdjCount
}

func countBlackTilesForDay(tiles map[Coordinate]int) (map[Coordinate]int, int) {
	nextState := make(map[Coordinate]int)
	for tile := range tiles {
		color := tiles[tile]
		// Any white tile with exactly 2 black tiles immediately adjacent to it is flipped to black.
		if color == 0 { //color is white
			adjBlack := countAdjacentBlackTiles(tile, tiles, nextState)
			if adjBlack == 2 {
				nextState[tile] ^= 1
			} else {
				nextState[tile] = tiles[tile]
			}
			// Any black tile with zero or more than 2 black tiles immediately adjacent to it is flipped to white.
		} else { // color is black
			adjBlack := countAdjacentBlackTiles(tile, tiles, nextState)
			if adjBlack == 0 || adjBlack > 2 {
				nextState[tile] ^= 1
			} else {
				nextState[tile] = tiles[tile]
			}
		}
	}

	res := 0
	for _, color := range nextState {
		res += color
	}
	return nextState, res
}

func partTwo() {
	numBlackTiles := 0
	tiles := getInitialColors()
	for i := 0; i < *numDays; i++ {
		nextState, count := countBlackTilesForDay(tiles)
		numBlackTiles += count
		tiles = nextState
	}
	fmt.Println(numBlackTiles)
}

type Coordinate struct {
	x, y, z int
}

func main() {
	partOne()
	partTwo()
}
