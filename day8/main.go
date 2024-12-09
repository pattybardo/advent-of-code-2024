package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var filename = "input.txt"

type coordinate struct {
	x, y int
}

type vector struct {
	x, y int
}

type bounds struct {
	x, y int
}

func main() {
	fmt.Println(solve())
}

func solve() int {
	coordinateChars := readInput()

	nodeMap := mapNodes(coordinateChars)

	bound := bounds{len(coordinateChars[0]), len(coordinateChars)}

	antinodes := determineAntinodes(nodeMap, bound)

	visualize(coordinateChars, antinodes)

	return len(antinodes)
}

func visualize(coordinateChars [][]rune, antinodes map[coordinate]struct{}) {
	for i, row := range coordinateChars {
		fmt.Printf("Row %2d: ", i)
		for j, r := range row {
			temp := coordinate{j, i}
			_, exists := antinodes[temp]
			if exists && r == '.' {
				fmt.Printf("%c ", '#')
			} else {
				fmt.Printf("%c ", r) // Print each rune as a character
			}
		}
		fmt.Println() // Newline after each row
	}
}

func determineAntinodes(nodeMap map[rune][]coordinate, bound bounds) map[coordinate]struct{} {
	antinodes := make(map[coordinate]struct{})
	for nodeType, coords := range nodeMap {
		addAntinodes(coords, antinodes, bound)
		fmt.Println(string(nodeType), coords)
	}
	return antinodes
}

func addAntinodes(coords []coordinate, antinodes map[coordinate]struct{}, bound bounds) {
	if len(coords) != 1 {
		for _, coord := range coords {
			antinodes[coord] = struct{}{}
		}
	}
	for i := 0; i < len(coords)-1; i++ {
		determineVectors(coords[i], coords[i+1:], antinodes, bound)
	}
}

func determineVectors(P coordinate, coordsToCheck []coordinate, antinodes map[coordinate]struct{}, bound bounds) {
	for _, Q := range coordsToCheck {
		diffVector := differenceVector(P, Q)
		addValidAntinodes(P, Q, diffVector, bound, antinodes)
	}
}

func addValidAntinodes(P, Q coordinate, diffVector vector, bound bounds, antinodes map[coordinate]struct{}) {
	done := false
	for i := 1; !done; i++ {
		firstAntinode := coordinate{Q.x + i*diffVector.x, Q.y + i*diffVector.y}
		secondAntinode := coordinate{P.x - i*diffVector.x, P.y - i*diffVector.y}

		firstValid := checkBounds(firstAntinode, bound)
		secondValid := checkBounds(secondAntinode, bound)

		if firstValid {
			antinodes[firstAntinode] = struct{}{}
		}
		if secondValid {
			antinodes[secondAntinode] = struct{}{}
		}
		if !secondValid && !firstValid {
			done = true
		}
	}
}

func checkBounds(antinode coordinate, bound bounds) bool {
	return antinode.x >= 0 && antinode.x < bound.x && antinode.y >= 0 && antinode.y < bound.y
}

func differenceVector(P, Q coordinate) vector {
	// Q-P. Therefore above we will Add this vector to Q, and subtract it from P
	// to get the antinodes
	return vector{Q.x - P.x, Q.y - P.y}
}

func mapNodes(coordinateChars [][]rune) map[rune][]coordinate {
	nodeMap := make(map[rune][]coordinate)
	// y
	for y, val := range coordinateChars {
		// x
		for x := range val {
			if coordinateChars[y][x] == '.' {
				continue
			}
			// I am flipping the coordinates now so it is x,y
			nodeMap[coordinateChars[y][x]] = append(nodeMap[coordinateChars[y][x]], coordinate{x, y})
		}
	}
	return nodeMap
}

func readInput() [][]rune {
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	text := string(contents)

	lines := strings.Split(strings.TrimSuffix(text, "\n"), "\n")

	var arrayOfArrays [][]rune
	for _, line := range lines {
		arrayOfArrays = append(arrayOfArrays, []rune(line))
	}

	// Split the text into lines
	return arrayOfArrays
}
