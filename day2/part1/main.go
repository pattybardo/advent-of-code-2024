package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type inputLine struct {
	x []int
}

func main() {
	fmt.Println(safety_calculation("input.txt"))
}

func safety_calculation(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var counter int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Read the line
		line := scanner.Text()

		// Split the line into string elements
		parts := strings.Fields(line) // Splits by whitespace

		// Convert the string elements to integers
		var nums []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				return 0
			}
			nums = append(nums, num)
		}

		// Append the integer array to the data slice

		counter += safety_check(nums)
	}

	// Check for scanner errors
	return counter
}

func safety_check(nums []int) int {
	// Rule 1: Either all increasing or all decreasing
	// Rule 2: At least a difference of 1 and max diff of 3

	rule_1 := nums[0] < nums[1] // Increasing = true

	for i := 1; i < len(nums); i++ {
		if rule_1 {
			if nums[i] < nums[i-1] {
				return 0
			}
			if nums[i]-nums[i-1] > 3 || nums[i]-nums[i-1] < 1 {
				return 0
			}
		} else {
			if nums[i] > nums[i-1] {
				return 0
			}
			if nums[i-1]-nums[i] > 3 || nums[i-1]-nums[i] < 1 {
				return 0
			}

		}
	}
	return 1
}
