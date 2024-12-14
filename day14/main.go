package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// var (
//
//	wide     = 11
//	tall     = 7
//	time     = 100
//	filename = "test.txt"
//
// )
var (
	wide = 101
	tall = 103
	// time     = 100
	filename = "input.txt"
)

type Position struct {
	X int
	Y int
}

type Velocity struct {
	X, Y int
}

type Robot struct {
	Position Position
	Velocity Velocity
}

func main() {
	fmt.Println(solve())
}

func solve() int {
	robots := readInput()

	// fmt.Println(robots)

	maxStraightLines := 0

	for time := 0; time < 10000; time++ {
		mapMatrix := initializeMatrix()
		for _, robot := range robots {
			pos := calculatePosition(robot, time)
			mapMatrix[pos.Y][pos.X] += 1
		}
		linesOverThreshold := calculateStraightLines(mapMatrix, 4)
		if linesOverThreshold > maxStraightLines {
			maxStraightLines = linesOverThreshold
			visualizeMatrix(mapMatrix)
			fmt.Println("--------------------")
			fmt.Println("T: ", time)
		}
	}

	// visualizeMatrix(mapMatrix)

	return 0
	// return calculateSafetyFactor(mapMatrix)
}

func calculateStraightLines(mapMatrix [][]int, threshold int) int {
	totalConsecutive := 0
	for i := range len(mapMatrix) {
		consecutive := 0
		for j := range len(mapMatrix[i]) - 1 {
			if mapMatrix[i][j] == 0 {
				consecutive = 0
				continue
			}
			if mapMatrix[i][j] == mapMatrix[i][j+1] {
				consecutive += 1
				if consecutive >= threshold {
					totalConsecutive += 1
					continue
				}
			} else {
				consecutive = 0
			}
		}
	}
	return totalConsecutive
}

func calculateSafetyFactor(mapMatrix [][]int) int {
	quadrants := []int{0, 0, 0, 0}
	mid_wide := (wide - 1) / 2
	mid_tall := (tall - 1) / 2
	// y-axis
	for i, row := range mapMatrix {
		// x-axis
		for j, entry := range row {
			if j < mid_wide {
				if i < mid_tall {
					quadrants[0] += entry
				} else if i > mid_tall {
					quadrants[3] += entry
				}
			} else if j > mid_wide {
				if i < mid_tall {
					quadrants[1] += entry
				} else if i > mid_tall {
					quadrants[2] += entry
				}
			}
		}
	}
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func calculatePosition(robot Robot, time int) Position {
	Ptx := (robot.Position.X + robot.Velocity.X*time) % wide
	Pty := (robot.Position.Y + robot.Velocity.Y*time) % tall

	if Ptx < 0 {
		Ptx += wide
	}
	if Pty < 0 {
		Pty += tall
	}
	return Position{Ptx, Pty}
}

// Function to initialize a matrix with a default value
func initializeMatrix() [][]int {
	matrix := make([][]int, tall)
	for i := range matrix {
		matrix[i] = make([]int, wide)
		for j := range matrix[i] {
			matrix[i][j] = 0
		}
	}
	return matrix
}

// Function to visualize a matrix
func visualizeMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, entry := range row {
			if entry == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(entry)
			}
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}

func readInput() []Robot {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	// Regular expression to match the input format
	lineRegex := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

	var robots []Robot

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		matches := lineRegex.FindStringSubmatch(line)
		if len(matches) != 5 {
			return nil
		}

		// Parse position and velocity values
		pX, _ := strconv.Atoi(matches[1])
		pY, _ := strconv.Atoi(matches[2])
		vX, _ := strconv.Atoi(matches[3])
		vY, _ := strconv.Atoi(matches[4])

		// Append to the robots slice
		robots = append(robots, Robot{
			Position: Position{X: pX, Y: pY},
			Velocity: Velocity{X: vX, Y: vY},
		})
	}

	if err := scanner.Err(); err != nil {
		return nil
	}

	return robots
}
