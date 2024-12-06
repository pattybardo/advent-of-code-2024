package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type position struct {
	x int
	y int
}

func main() {
	fmt.Println(calculate_xmas("test.txt"))
}

func calculate_xmas(filename string) int {
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	text := string(contents)

	// Split the text into lines
	lines := splitIntoLines(text)

	// Convert each line into a slice of characters (array of runes)
	var arrayOfArrays [][]rune
	for _, line := range lines {
		arrayOfArrays = append(arrayOfArrays, []rune(line))
	}

	counter := 0

	for i := 1; i < len(arrayOfArrays)-1; i++ {
		for j := 1; j < len(arrayOfArrays[0])-1; j++ {
			if arrayOfArrays[i][j] == 'A' {
				counter += check_cross(position{i, j}, arrayOfArrays)
			}
		}
	}

	return counter
}

// Helper function to split a string into lines
func splitIntoLines(text string) []string {
	return strings.Split(strings.TrimSuffix(text, "\n"), "\n")
}

func check_cross(position_of_a position, input [][]rune) int {
	var my_map map[rune]int
	my_map = make(map[rune]int)
	my_map[input[position_of_a.x-1][position_of_a.y-1]] += 1
	my_map[input[position_of_a.x+1][position_of_a.y-1]] += 1
	my_map[input[position_of_a.x-1][position_of_a.y+1]] += 1
	my_map[input[position_of_a.x+1][position_of_a.y+1]] += 1

	if my_map['M'] == 2 && my_map['S'] == 2 && input[position_of_a.x-1][position_of_a.y-1] != input[position_of_a.x+1][position_of_a.y+1] {
		return 1
	} else {
		return 0
	}
}
