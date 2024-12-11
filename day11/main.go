package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	filename = "input.txt"
	limit    = 75
)

type myTuple struct {
	val   string
	depth int
}

func main() {
	fmt.Println(solve())
}

func solve() int {
	input := readInput()

	calculatedNodes := make(map[myTuple]int)

	output := blinkX(input, calculatedNodes)
	fmt.Println(output)

	return 0
}

func blinkX(input []string, calculatedNodes map[myTuple]int) int {
	counter := 0
	// Recurse
	for _, val := range input {
		counter += blink(val, calculatedNodes, 0)
	}

	return counter
}

func blink(val string, calculatedNodes map[myTuple]int, depth int) int {
	if depth >= limit {
		// fmt.Println(val)
		return 1
	} else {
		counter, exists := calculatedNodes[myTuple{val, depth}]
		if exists {
			// fmt.Println("EXISTS: ", val)
			return counter
		} else {
			if val == "0" {
				counter += blink("1", calculatedNodes, depth+1)
			} else if len(val)%2 == 0 {
				// fmt.Println("DEPTH: ", depth)
				x, y := splitDigits(val)
				// fmt.Println(x, y)
				counter += blink(x, calculatedNodes, depth+1)
				counter += blink(y, calculatedNodes, depth+1)
			} else {
				counter += blink(recalculate(val), calculatedNodes, depth+1)
			}
			calculatedNodes[myTuple{val, depth}] = counter
		}
		return counter
	}
}

func next(val string) {
}

func recalculate(val string) string {
	num, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return strconv.Itoa(num * 2024)
}

func splitDigits(val string) (string, string) {
	x := val[:len(val)/2]
	y := val[len(val)/2:]
	if strings.TrimLeft(y, "0") == "" {
		y = "0"
	} else {
		y = strings.TrimLeft(y, "0")
	}
	return x, y
}

func readInput() []string {
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(strings.TrimRight(string(contents), "\n"), " ")

	return data
}
