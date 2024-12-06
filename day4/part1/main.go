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
	fmt.Println(calculate_xmas("input.txt"))
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

	target := []rune{'X', 'M', 'A', 'S'}

	counter := 0

	for i, arr_val := range arrayOfArrays {
		for j, val := range arr_val {
			if val == 'X' {
				counter += check_diagonals(position{i, j}, arrayOfArrays, target)
				counter += check_verticals(position{i, j}, arrayOfArrays, target)
				counter += check_horizontals(position{i, j}, arrayOfArrays, target)
			}
		}
	}

	return counter
}

// Helper function to split a string into lines
func splitIntoLines(text string) []string {
	return strings.Split(strings.TrimSuffix(text, "\n"), "\n")
}

func check_diagonals(position_of_x position, input [][]rune, target []rune) int {
	counter := 0

	if position_of_x.x-3 >= 0 && position_of_x.y-3 >= 0 {
		should_count := true
		for i := 1; i < 4; i++ {
			if input[position_of_x.x-i][position_of_x.y-i] != target[i] {
				should_count = false
			}
		}
		if should_count {
			counter += 1
		}
	}

	if position_of_x.x+3 < len(input[0]) && position_of_x.y-3 >= 0 {
		should_count := true
		for i := 1; i < 4; i++ {
			if input[position_of_x.x+i][position_of_x.y-i] != target[i] {
				should_count = false
			}
		}
		if should_count {
			counter += 1
		}
	}

	if position_of_x.x-3 >= 0 && position_of_x.y+3 < len(input) {
		should_count := true
		for i := 1; i < 4; i++ {
			if input[position_of_x.x-i][position_of_x.y+i] != target[i] {
				should_count = false
			}
		}
		if should_count {
			counter += 1
		}
	}

	if position_of_x.x+3 < len(input[0]) && position_of_x.y+3 < len(input) {
		should_count := true
		for i := 1; i < 4; i++ {
			if input[position_of_x.x+i][position_of_x.y+i] != target[i] {
				should_count = false
			}
		}
		if should_count {
			counter += 1
		}
	}

	fmt.Println(counter)

	return counter
}

func check_horizontals(position_of_x position, input [][]rune, target []rune) int {
	counter := 0

	if position_of_x.x-3 >= 0 {
		should_count := true
		for i := 1; i < 4; i++ {
			if input[position_of_x.x-i][position_of_x.y] != target[i] {
				should_count = false
			}
		}
		if should_count {
			counter += 1
		}
	}

	if position_of_x.x+3 < len(input[0]) {
		should_count := true
		for i := 1; i < 4; i++ {
			if input[position_of_x.x+i][position_of_x.y] != target[i] {
				should_count = false
			}
		}
		if should_count {
			counter += 1
		}
	}

	return counter
}

func check_verticals(position_of_x position, input [][]rune, target []rune) int {
	counter := 0

	if position_of_x.y-3 >= 0 {
		should_count := true
		for i := 1; i < 4; i++ {
			if input[position_of_x.x][position_of_x.y-i] != target[i] {
				should_count = false
			}
		}
		if should_count {
			counter += 1
		}
	}

	if position_of_x.y+3 < len(input[0]) {
		should_count := true
		for i := 1; i < 4; i++ {
			if input[position_of_x.x][position_of_x.y+i] != target[i] {
				should_count = false
			}
		}
		if should_count {
			counter += 1
		}
	}

	return counter
}
