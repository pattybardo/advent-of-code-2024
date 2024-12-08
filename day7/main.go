package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var filename = "test.txt"

func main() {
	fmt.Println(solve())
}

func solve() int {
	lines := readInput()

	combinations := generateCombinations(12, 3)
	counter := 0
	for _, line := range lines {
		counter += checkInputLine(combinations, line)
	}

	// fmt.Println(combinations)

	return counter
}

func checkInputLine(combinations map[int][][]int, line string) int {
	finalValue, inputs := parseLine(line)
	for _, combination := range combinations[len(inputs)-1] {
		if checkCombination(combination, inputs) == finalValue {
			return finalValue
		}
	}
	return 0
}

func checkCombination(combination []int, inputs []int) int {
	counter := inputs[0]
	for i, operation := range combination {
		counter = evaluator(operation, counter, inputs[i+1])
	}
	return counter
}

func evaluator(operation int, x, y int) int {
	switch {
	case operation == 0:
		return x + y
	case operation == 1:
		return x * y
	case operation == 2:
		return concatenateIntegers(x, y)
	}
	return 0
}

func concatenateIntegers(a, b int) int {
	// Convert integers to strings
	strA := strconv.Itoa(a)
	strB := strconv.Itoa(b)

	// Concatenate the strings
	concatenated := strA + strB

	// Convert the concatenated string back to an integer
	result, _ := strconv.Atoi(concatenated)

	return result
}

func parseLine(line string) (int, []int) {
	// 1234: 1 2 3 4
	str := strings.Split(line, ":")
	finalValue, err := strconv.Atoi(str[0]) // Convert string to integer
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return -1, []int{-1}
	}
	numStrings := strings.Split(strings.TrimSpace(str[1]), " ")
	var inputs []int
	for _, num := range numStrings {
		parsedNum, err := strconv.Atoi(num) // Convert string to integer
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return -1, []int{-1}
		}
		inputs = append(inputs, parsedNum)
	}
	return finalValue, inputs
}

func generateCombinations(maxLength int, numOps int) map[int][][]int {
	combinationsMap := make(map[int][][]int)

	// Helper function to build combinations using recursion
	var generate func([]int, int, int) [][]int
	generate = func(current []int, depth, length int) [][]int {
		if depth == length {
			combination := make([]int, len(current))
			copy(combination, current)
			return [][]int{combination}
		}

		results := [][]int{}
		for i := 0; i < numOps; i++ {
			results = append(results, generate(append(current, i), depth+1, length)...)
		}
		return results
	}

	// Generate combinations for each possible length up to maxLength
	for length := 1; length <= maxLength; length++ {
		combinationsMap[length] = generate([]int{}, 0, length)
	}

	return combinationsMap
}

func readInput() []string {
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	text := string(contents)

	// Split the text into lines
	return splitIntoLines(text)
}

func splitIntoLines(text string) []string {
	return strings.Split(strings.TrimSuffix(text, "\n"), "\n")
}
