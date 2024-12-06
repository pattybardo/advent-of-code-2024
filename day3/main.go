package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type inputLine struct {
	x []int
}

func main() {
	fmt.Println(mul_calculation("test.txt"))
}

func mul_calculation(filename string) int {
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	text := string(contents)

	pattern := `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`
	re := regexp.MustCompile(pattern)

	// Find all matches
	matches := re.FindAllStringSubmatch(text, -1)

	counter := 0
	do := true

	for _, match := range matches {
		if match[0] == "do()" {
			do = true
		} else if match[0] == "don't()" {
			do = false
		} else if len(match) > 2 && do {
			num1, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Println(match)
				fmt.Println("Error converting to integer:", err)
				return 0
			}
			num2, err := strconv.Atoi(match[2])
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				return 0
			}
			counter += num1 * num2
		}

		fmt.Println(match)
	}
	return counter
}
