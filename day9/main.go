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
	positions := splitInput(input)
	printPositions(positions, input[0])
	// fmt.Println(positions)

	// fmt.Println(arrangedPositions)
	output := calculateOutput(positions)

	return output
}

func printPositions(positions []int, start int) {
	for i := 0; i < len(positions); i++ {
		if positions[i] == 0 && i >= start {
			fmt.Print(".")
		} else {
			fmt.Print(positions[i])
		}
	}
	fmt.Print("\n")
}

func calculateOutput(arrangedPositions []int) int {
	counter := 0
	for i, val := range arrangedPositions {
		counter += i * val
	}
	return counter
}

func splitInput(input []int) []int {
	counter := 0
	odd := true
	for i := range len(input) {
		if odd {
			counter += input[i]
		}
	}

	positions := make([]int, counter)

	updatedInput := make([]int, len(input))
	copy(updatedInput, input)
	leftoverArray := make([]int, len(input))

	rightSwitcher := true

	for right := len(input) - 1; right >= 0; right-- {
		counter = 0
		odd = false
		has_moved := false
		if rightSwitcher {
			for left := range len(input) {
				if left >= right {
					break
				}
				if odd {
					if updatedInput[left] >= input[right] {
						updatedInput[left] -= input[right]
						for i := range input[right] {
							positions[counter+leftoverArray[left]+i] = right / 2
						}
						if updatedInput[left] > 0 {
							leftoverArray[left] += input[right]
						}
						has_moved = true
						break
					}
				}
				counter += input[left]
				odd = !odd
			}
			if has_moved == false {
				for i := range input[right] {
					positions[counter+i] = right / 2
				}
			}
		}
		rightSwitcher = !rightSwitcher
	}

	return positions
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
