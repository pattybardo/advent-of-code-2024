package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var filename = "input.txt"

func main() {
	fmt.Println(solve())
}

func solve() int {
	coordinateChars := readInput()

	for i, row := range coordinateChars {
		fmt.Printf("Row %2d: ", i)
		for _, r := range row {
			fmt.Printf("%c ", r) // Print each rune as a character
		}
		fmt.Println() // Newline after each row
	}

	counter := 0

	return counter
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
