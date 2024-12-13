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

type direction struct {
	x, y int
}

type gardenNode struct {
	coordinate     coordinate
	perimeterCount int
}

var directions = []direction{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func main() {
	fmt.Println(solve())
}

func solve() int {
	coordinateChars := readInput()

	visualize(coordinateChars)

	// Map from the first gardenNode in a group to the other nodes belonging to that group
	gardenGrouping := make(map[coordinate][]gardenNode)
	// Map to check if we have visited a node already
	visitedNodes := make(map[coordinate]struct{})

	gardenBounds := len(coordinateChars)
	// coordinateChars is always square
	for y := range gardenBounds {
		for x := range gardenBounds {
			_, exists := visitedNodes[coordinate{x, y}]
			if exists {
				continue
			}
			processGarden(coordinateChars, visitedNodes, gardenGrouping, coordinate{x, y}, coordinate{x, y}, gardenBounds)
		}
	}
	visualizeGrouping(coordinateChars, gardenGrouping)

	return 0
}

func processGarden(input [][]rune, visitedNodes map[coordinate]struct{}, gardenGrouping map[coordinate][]gardenNode, parentCoordinate, nodeCoordinate coordinate, bounds int) {
	perimeterCounter := 0
	visitedNodes[nodeCoordinate] = struct{}{}
	for _, direction := range directions {
		newCoordinate := coordinate{nodeCoordinate.x + direction.x, nodeCoordinate.y + direction.y}
		if !checkBounds(input, newCoordinate, nodeCoordinate, bounds) {
			perimeterCounter += 1
			continue
		}
		_, exists := visitedNodes[newCoordinate]
		if exists {
			continue
		}
		processGarden(input, visitedNodes, gardenGrouping, parentCoordinate, newCoordinate, bounds)
	}
	_, exists := gardenGrouping[parentCoordinate]
	if !exists {
		gardenGrouping[parentCoordinate] = []gardenNode{{nodeCoordinate, perimeterCounter}}
	} else {
		gardenGrouping[parentCoordinate] = append(gardenGrouping[parentCoordinate], gardenNode{nodeCoordinate, perimeterCounter})
	}
}

func visualize(coordinateChars [][]rune) {
	for i, row := range coordinateChars {
		fmt.Printf("Row %2d: ", i)
		for _, r := range row {
			fmt.Printf("%c ", r) // Print each rune as a character
		}
		fmt.Println() // Newline after each row
	}
}

func visualizeGrouping(coordinateChars [][]rune, gardenGrouping map[coordinate][]gardenNode) {
	counter := 0
	for _, valArr := range gardenGrouping {
		areaCounter := 0
		perimeterCounter := 0
		// fmt.Printf("%c: ", coordinateChars[key.y][key.x])
		for _, val := range valArr {
			// fmt.Print(val, ", ")
			areaCounter += 1
			perimeterCounter += val.perimeterCount
		}
		counter += areaCounter * perimeterCounter
		// fmt.Print("\n")
	}
	fmt.Println(counter)
}

func checkBounds(input [][]rune, newCoordinate, coordinate2 coordinate, bound int) bool {
	if newCoordinate.x >= 0 && newCoordinate.x < bound && newCoordinate.y >= 0 && newCoordinate.y < bound {
		if input[newCoordinate.y][newCoordinate.x] != input[coordinate2.y][coordinate2.x] {
			return false
		}
		return true
	}
	return false
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
