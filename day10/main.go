package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var filename = "input.txt"

type position struct {
	x, y int
}
type vector struct {
	x, y int
}

func main() {
	input := readInput()
	fmt.Println("Slice of slices:")
	for _, row := range input {
		fmt.Println(row)
	}
	endPositions := make(map[position]map[position]int)
	for y := range input {
		for x := range input[0] {
			if input[y][x] == 0 {
				trailRecursion(position{x, y}, position{x, y}, 1, input, endPositions)
			}
		}
	}
	counter := 0
	for _, val := range endPositions {
		for _, num := range val {
			counter += num
		}
	}
	fmt.Println(counter)
}

func trailRecursion(startingPos, pos position, next int, input [][]int, output map[position]map[position]int) {
	if input[pos.y][pos.x] == 9 {
		if output[startingPos] == nil {
			output[startingPos] = make(map[position]int)
		}

		output[startingPos][pos] += 1
	}
	directions := []vector{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	for _, direction := range directions {
		if inbounds(pos, direction, len(input[0]), len(input)) {
			newPos := position{pos.x + direction.x, pos.y + direction.y}
			if input[newPos.y][newPos.x] == next {
				trailRecursion(startingPos, newPos, next+1, input, output)
			}
		}
	}
}

func inbounds(pos position, direction vector, x_bound int, y_bound int) bool {
	return pos.x+direction.x >= 0 && pos.x+direction.x < x_bound && pos.y+direction.y >= 0 && pos.y+direction.y < y_bound
}

func readInput() [][]int {
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	text := string(contents)
	// Split input into lines
	lines := strings.Split(strings.TrimSuffix(text, "\n"), "\n")

	var data [][]int

	for _, line := range lines {

		var row []int
		for _, char := range line {

			num, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Println("Error converting character to integer:", err)
				return nil
			}
			row = append(row, num)
		}
		data = append(data, row)
	}

	return data
}
