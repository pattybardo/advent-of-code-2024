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
	// If we can remove 1 entry and it still work, we gucci

	dampener_count := 0
	rule_1, should_continue := set_rule_1(nums) // Increasing = true
	if !should_continue {
		return 0
	}

	for i := 1 + dampener_count; i < len(nums); i++ {
		shouldReturn := checkRule(rule_1, nums[i-1], nums[i])
		if shouldReturn {
			if dampener_count == 0 {
				if i == len(nums)-1 {
					return 1
				}
				shouldReturn = checkRule(rule_1, nums[i-1], nums[i+1])
				if shouldReturn {

					shouldReturn = checkRule(rule_1, nums[i], nums[i+1])
					if shouldReturn {
						return 0
					} else {
						dampener_count += 1
						continue
					}
				} else {
					dampener_count += 1
					i++
				}
			} else {
				return 0
			}
		}
	}
	return 1
}

func set_rule_1(nums []int) (bool, bool) {
	counter := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			counter += 1
		} else {
			counter -= 1
		}
	}
	if counter > 0 {
		return true, true
	} else if counter < 0 {
		return false, true
	} else {
		return false, false
	}
}

func checkRule(rule_1 bool, num1 int, num2 int) bool {
	if rule_1 {
		if num2 < num1 {
			return true
		}
		if num2-num1 > 3 || num2-num1 < 1 {
			return true
		}
	} else {
		if num2 > num1 {
			return true
		}
		if num1-num2 > 3 || num1-num2 < 1 {
			return true
		}

	}
	return false
}
