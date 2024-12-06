package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	page_ordering_rules, pages_to_produce := splitRules(text)

	data := make(map[string][]string)
	for _, page_order := range page_ordering_rules {
		temp := strings.Split(page_order, "|")
		data[temp[0]] = append(data[temp[0]], temp[1])
	}

	counter := 0

	fmt.Println(data)
	for _, page_to_produce := range pages_to_produce {
		counter += calculate_pages(page_to_produce, data)
	}

	return counter
}

func calculate_pages(set_of_pages string, data map[string][]string) int {
	pages := strings.Split(set_of_pages, ",")

	is_valid := true

	for i := 1; i < len(pages); i++ {
		temp := is_in_valid_position(pages[:i], data[pages[i]])
		if temp == -1 {
			continue
		} else {

			elementToMove := pages[i]
			// Step 2: Remove the element from its current position
			pages = append(pages[:i], pages[i+1:]...)

			// Step 3: Insert the element at the target position
			pages = append(pages[:temp], append([]string{elementToMove}, pages[temp:]...)...)
			i = 1
			is_valid = false

		}
	}

	if !is_valid {
		num, err := strconv.Atoi(pages[(len(pages)-1)/2])
		if err != nil {
			fmt.Println("Error converting to integer:", err)
			return 0
		}
		return num
	}

	return 0
}

func is_in_valid_position(values_to_check []string, data_slice []string) int {
	for i, x := range values_to_check {
		for _, v := range data_slice {
			if v == x {
				return i
			}
		}
	}

	return -1
}

// Helper function to split a string into lines
func splitRules(text string) ([]string, []string) {
	temp := strings.Split(strings.TrimSuffix(text, "\n"), "\n\n")
	return strings.Split(temp[0], "\n"), strings.Split(temp[1], "\n")
}
