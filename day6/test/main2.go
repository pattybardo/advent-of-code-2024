package main

// package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Position struct {
	X, Y int
}

func main() {
	fmt.Println(guard_walk("input.txt"))
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

	x_dir := -1
	y_dir := -1

	move_in_x := false
	done := false

	adjacency_map := make(map[Position]map[Position]struct{})
	obstacles := make(map[Position]struct{})

	// return 0
	// my_set[pos] = struct{}{}
	for {
		step(&pos, &x_dir, &y_dir, &move_in_x, &done, arrayOfArrays, adjacency_map, obstacles, true)
		// my_set[pos] = struct{}{}
		if done {
			break
		}
	}

	//for i, val := range adjacency_map {
	//	fmt.Println(i, ":", val)
	//}

	for i := range obstacles {
		fmt.Println(i)
	}
	return len(obstacles)
}

func step(pos *Position, x_dir *int, y_dir *int, move_in_x *bool, done *bool, input [][]rune, adjacency_map map[Position]map[Position]struct{}, obstacles map[Position]struct{}, hack bool) {
	old_pos := Position{pos.X, pos.Y}
	if *move_in_x {
		if input[pos.Y][pos.X+*x_dir] == '#' {
			turn(x_dir, y_dir, move_in_x)
			if hack && !*done {
				if is_blocking_obstacle(Position{pos.X, pos.Y}, adjacency_map, *move_in_x, *x_dir, *y_dir, input) {
					obstacles[Position{pos.X, pos.Y + *y_dir}] = struct{}{}
				}
			}
			return
		}
		if hack && !*done {
			if is_blocking_obstacle(Position{pos.X, pos.Y}, adjacency_map, *move_in_x, *x_dir, *y_dir, input) {
				obstacles[Position{pos.X + *x_dir, pos.Y}] = struct{}{}
			}
		}
		pos.X += *x_dir
		*done = !should_continue(pos.X+*x_dir, pos.Y, input)

	} else {
		if input[pos.Y+*y_dir][pos.X] == '#' {
			turn(x_dir, y_dir, move_in_x)
			if hack && !*done {
				if is_blocking_obstacle(Position{pos.X, pos.Y}, adjacency_map, *move_in_x, *x_dir, *y_dir, input) {
					obstacles[Position{pos.X + *x_dir, pos.Y}] = struct{}{}
				}
			}
			return
		}
		if hack && !*done {
			if is_blocking_obstacle(Position{pos.X, pos.Y}, adjacency_map, *move_in_x, *x_dir, *y_dir, input) {
				obstacles[Position{pos.X, pos.Y + *y_dir}] = struct{}{}
			}
		}
		pos.Y += *y_dir

		*done = !should_continue(pos.X, pos.Y+*y_dir, input)

	}

	temp := Position{pos.X, pos.Y}

	if adjacency_map[temp] == nil {
		adjacency_map[temp] = make(map[Position]struct{})
	}

	//if !hack {
	//	fmt.Println("Simulating stepping from: ", old_pos, " to: ", temp)
	//} else {
	//	fmt.Println("Stepping from: ", old_pos, " to: ", temp)
	//}
	adjacency_map[temp][old_pos] = struct{}{}
}

func is_blocking_obstacle(old_pos Position, adjacency_map map[Position]map[Position]struct{}, move_in_x bool, x_dir, y_dir int, input [][]rune) bool {
	turn(&x_dir, &y_dir, &move_in_x)
	done := false
	exists := false
	copy_exists := false
	blocking_adjacency_map := deepCopyAdjacencyMap(adjacency_map)
	obstacles := make(map[Position]struct{})
	if move_in_x {
		_, exists = adjacency_map[Position{old_pos.X + x_dir, old_pos.Y}][old_pos]
	} else {
		_, exists = adjacency_map[Position{old_pos.X, old_pos.Y + y_dir}][old_pos]
	}
	if !exists {
		for {
			step(&old_pos, &x_dir, &y_dir, &move_in_x, &done, input, blocking_adjacency_map, obstacles, false)
			if move_in_x {
				_, exists = adjacency_map[Position{old_pos.X + x_dir, old_pos.Y}][old_pos]
				_, copy_exists = blocking_adjacency_map[Position{old_pos.X + x_dir, old_pos.Y}][old_pos]
			} else {
				_, exists = adjacency_map[Position{old_pos.X, old_pos.Y + y_dir}][old_pos]
				_, copy_exists = blocking_adjacency_map[Position{old_pos.X, old_pos.Y + y_dir}][old_pos]
			}
			if done {
				return false
			} else if exists {
				return true
			} else if copy_exists {
				return false
			}
		}
	}
	return true
}

func should_continue(x int, y int, input [][]rune) bool {
	if !check_bounds(x, y, input) {
		return false
	}
	return true
}

func check_bounds(x, y int, input [][]rune) bool {
	return x < len(input[0]) && x >= 0 && y >= 0 && y < len(input)
}

func turn(x_dir *int, y_dir *int, move_in_x *bool) {
	if *x_dir == -1 && *y_dir == -1 {
		*x_dir = 1
	} else if *x_dir == 1 && *y_dir == -1 {
		*y_dir = 1
	} else if *x_dir == 1 && *y_dir == 1 {
		*x_dir = -1
	} else if *x_dir == -1 && *y_dir == 1 {
		*y_dir = -1
	}
	*move_in_x = !*move_in_x
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

func deepCopyAdjacencyMap(original map[Position]map[Position]struct{}) map[Position]map[Position]struct{} {
	// Create a new map for the copy
	copy := make(map[Position]map[Position]struct{})

	// Iterate over the outer map
	for key, innerMap := range original {
		// Create a new inner map
		innerCopy := make(map[Position]struct{})

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
