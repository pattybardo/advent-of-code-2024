package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type inputs struct {
	first  []int
	second []int
}

func main() {
	parsed_inputs := parse_location_list("test.txt")
	sorted_inputs := sort_inputs(parsed_inputs)
	fmt.Println(calculate_distance(sorted_inputs))
	second_map := mapify_second(sorted_inputs.second)
	fmt.Println(calculate_similarity(sorted_inputs.first, second_map))
}

func calculate_similarity(first_list []int, second_map map[int]int) int {
	counter := 0
	for _, val := range first_list {
		if second_map[val] != 0 {
			counter += val * second_map[val]
		}
	}
	return counter
}

func mapify_second(inputs []int) map[int]int {
	m := make(map[int]int)
	for _, val := range inputs {
		if m[val] == 0 {
			m[val] = 1
		} else {
			m[val] = m[val] + 1
		}
	}
	return m
}

func calculate_distance(sorted_inputs inputs) int {
	var counter int
	for i, val := range sorted_inputs.first {
		temp := val - sorted_inputs.second[i]
		if temp < 0 {
			counter += -temp
		} else {
			counter += temp
		}
	}
	return counter
}

func parse_location_list(filename string) inputs {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	var first_column []int
	var second_column []int

	ticker := true
	for scanner.Scan() {
		token := scanner.Text()
		num, err := strconv.Atoi(token)
		if err != nil {
			fmt.Println(err)
		}
		if ticker {
			first_column = append(first_column, num)
			ticker = false
		} else {
			second_column = append(second_column, num)
			ticker = true
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return inputs{first_column, second_column}
}

func sort_inputs(unsorted_inputs inputs) inputs {
	sort.Slice(unsorted_inputs.first, func(i, j int) bool {
		return unsorted_inputs.first[i] < unsorted_inputs.first[j]
	})

	sort.Slice(unsorted_inputs.second, func(i, j int) bool {
		return unsorted_inputs.second[i] < unsorted_inputs.second[j]
	})

	return unsorted_inputs
}
