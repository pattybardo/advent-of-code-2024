package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var filename = "input.txt"

func main() {
	fmt.Println(solve())
}

func solve() int {
	input := readInput()

	// fmt.Println(input)
	positions, emptySpots := splitInput(input)

	// fmt.Println(positions)
	arrangedPositions := moveFileBlocks(positions, emptySpots)

	fmt.Println(arrangedPositions)
	output := calculateOutput(arrangedPositions)

	return output
}

func calculateOutput(arrangedPositions []int) int {
	counter := 0
	for i, val := range arrangedPositions {
		counter += i * val
	}
	return counter
}

func moveFileBlocks(positions []int, emptySpots []int) []int {
	counter := len(positions) - 1
	emptyCounter := 0
	for range emptySpots {
		if counter <= emptyCounter {
			break
		} else if positions[counter] == 0 {
			counter--
		} else {
			positions[emptySpots[emptyCounter]] = positions[counter]
			positions[counter] = 0
			counter--
			emptyCounter++
		}
	}
	return positions
}

func splitInput(input []int) ([]int, []int) {
	var emptySpots []int

	counter := 0
	odd := true
	for i := range len(input) {
		if odd {
			counter += input[i]
		}
	}

	positions := make([]int, counter)

	counter = 0
	odd = true
	id_counter := 0

	for i := range len(input) {
		if odd {
			for range input[i] {
				positions[counter] = id_counter
				counter++
			}
			id_counter++
		} else {
			for range input[i] {
				emptySpots = append(emptySpots, counter)
				counter++
			}
		}
		odd = !odd
	}

	return positions, emptySpots
}

func readInput() []int {
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	data := string(contents)

	var result []int

	for _, char := range strings.TrimSpace(data) {
		// Convert rune to string, then to int
		num, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Println(err)
			return nil // Return an error if conversion fails
		}
		result = append(result, num)
	}

	return result
}
