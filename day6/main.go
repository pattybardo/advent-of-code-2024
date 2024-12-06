package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Position struct {
	X, Y int
}

type Direction struct {
	x_dir, y_dir int
}

func main() {
	fmt.Println(guard_walk("test2.txt"))
}

func guard_walk(filename string) int {
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

	y, x := findGuardStartPosition(arrayOfArrays)

	pos := Position{x, y}

	dir := Direction{0, -1}
	done := false

	adjacency_map := make(map[Position]map[Direction]struct{})
	obj_map := make(map[Position]struct{})

	for !done {
		simulate(pos, dir, arrayOfArrays, adjacency_map, obj_map)
		step(&pos, &dir, &done, arrayOfArrays, adjacency_map)
	}

	for val := range adjacency_map {
		fmt.Println(val)
	}
	fmt.Println("---")
	for val := range obj_map {
		fmt.Println(val)
	}

	fmt.Println(len(adjacency_map))
	return len(obj_map)
}

func simulate(pos Position, dir Direction, input [][]rune, adjacency_map map[Position]map[Direction]struct{}, obj_map map[Position]struct{}) {
	obj_pos := Position{pos.X + dir.x_dir, pos.Y + dir.y_dir}
	if input[obj_pos.Y][obj_pos.X] == '#' {
		return
	}
	// Deep Copies
	input_copy := deepCopyInput(input)
	adj_map_copy := deepCopyAdjacencyMap(adjacency_map)

	input_copy[pos.Y+dir.y_dir][pos.X+dir.x_dir] = '#'
	turn(&dir)

	cp_done := false

	for !cp_done {
		is_loop := step(&pos, &dir, &cp_done, input_copy, adj_map_copy)
		if is_loop {
			obj_map[obj_pos] = struct{}{}
			return
		}
	}
}

func step(pos *Position, dir *Direction, done *bool, input [][]rune, adjacency_map map[Position]map[Direction]struct{}) bool {
	changeDirection(input, pos, dir)

	if adjacency_map[*pos] == nil {
		adjacency_map[*pos] = make(map[Direction]struct{})
	}
	// If this entry already exists, then we are in a loop
	_, is_loop := adjacency_map[*pos][*dir]
	adjacency_map[*pos][*dir] = struct{}{}

	// fmt.Println("Updating Map, ", *pos, Position{pos.X + dir.x_dir, pos.Y + dir.y_dir})

	updatePosition(pos, dir)

	*done = !within_bounds(*pos, *dir, input)
	return is_loop
}

func updatePosition(pos *Position, dir *Direction) {
	pos.X += dir.x_dir
	pos.Y += dir.y_dir
}

func changeDirection(input [][]rune, pos *Position, dir *Direction) {
	for {
		if input[pos.Y+dir.y_dir][pos.X+dir.x_dir] == '#' {
			turn(dir)
		} else {
			break
		}
	}
}

func within_bounds(pos Position, dir Direction, input [][]rune) bool {
	x := pos.X + dir.x_dir
	y := pos.Y + dir.y_dir
	return x < len(input[0]) && x >= 0 && y >= 0 && y < len(input)
}

func turn(dir *Direction) {
	// fmt.Println("TURNING")
	if dir.x_dir == 0 && dir.y_dir == -1 {
		dir.x_dir = 1
		dir.y_dir = 0
	} else if dir.x_dir == 1 && dir.y_dir == 0 {
		dir.y_dir = 1
		dir.x_dir = 0
	} else if dir.x_dir == 0 && dir.y_dir == 1 {
		dir.x_dir = -1
		dir.y_dir = 0
	} else if dir.x_dir == -1 && dir.y_dir == 0 {
		dir.y_dir = -1
		dir.x_dir = 0
	}
}

func findGuardStartPosition(input [][]rune) (int, int) {
	for i, val := range input {
		for j, pos := range val {
			if pos == '^' {
				return i, j
			}
		}
	}
	return -1, -1
}

func deepCopyInput(input [][]rune) [][]rune {
	copyInput := make([][]rune, len(input))
	for i := range input {
		copyInput[i] = append([]rune{}, input[i]...)
	}
	return copyInput
}

func deepCopyAdjacencyMap(original map[Position]map[Direction]struct{}) map[Position]map[Direction]struct{} {
	// Create a new map for the copy
	copy := make(map[Position]map[Direction]struct{})

	// Iterate over the outer map
	for key, innerMap := range original {
		// Create a new inner map
		innerCopy := make(map[Direction]struct{})

		// Copy each entry from the inner map
		for innerKey := range innerMap {
			innerCopy[innerKey] = struct{}{}
		}

		// Add the copied inner map to the outer map
		copy[key] = innerCopy
	}

	return copy
}

// Helper function to split a string into lines
func splitIntoLines(text string) []string {
	return strings.Split(strings.TrimSuffix(text, "\n"), "\n")
}
