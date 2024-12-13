package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var filename = "../input.txt"

type Position struct {
	X int
	Y int
}

type Button struct {
	Name     string
	Position Position
}

type Prize struct {
	Position Position
}

type Entry struct {
	A     Button
	B     Button
	Prize Prize
}

func main() {
	fmt.Println(solve())
}

func solve() int {
	entries := readInput()

	counter := 0
	for _, entry := range entries {
		counter += calculate(entry)
	}

	return counter
}

func calculate(entry Entry) int {
	var possibleTokens []int
	a_tokens := 3
	b_token := 1
	for a := range 101 {
		val_x := entry.Prize.Position.X - entry.A.Position.X*a
		val_y := entry.Prize.Position.Y - entry.A.Position.Y*a
		if val_x%entry.B.Position.X == 0 && val_y%entry.B.Position.Y == 0 {
			if val_x/entry.B.Position.X == val_y/entry.B.Position.Y {
				b := val_x / entry.B.Position.X
				// fmt.Println("A: ", a, " B: ", b)
				if b > 0 {
					possibleTokens = append(possibleTokens, a_tokens*a+b*b_token)
				}
			}
		}
	}
	if len(possibleTokens) > 0 {
		fmt.Println((possibleTokens))
		return slices.Min(possibleTokens)
	}
	// fmt.Println("Not Possible")
	return 0
}

func readInput() []Entry {
	var entries []Entry
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(string(contents), "\n\n")

	for _, input := range inputs {

		buttons, prize, err := parseInput(input)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		for _, button := range buttons {
			fmt.Printf("Button %s: X=%d, Y=%d\n", button.Name, button.Position.X, button.Position.Y)
		}
		fmt.Printf("Prize: X=%d, Y=%d\n", prize.Position.X, prize.Position.Y)
		entries = append(entries, Entry{buttons[0], buttons[1], prize})
	}

	return entries
}

func parseInput(input string) ([]Button, Prize, error) {
	// Split the input into lines
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Regular expressions for parsing
	buttonRegex := regexp.MustCompile(`Button (\w+): X([+\-=])(\d+), Y([+\-=])(\d+)`)
	prizeRegex := regexp.MustCompile(`Prize: X([+\-=])(\d+), Y([+\-=])(\d+)`)

	var buttons []Button
	var prize Prize

	for _, line := range lines {
		if strings.HasPrefix(line, "Button") {
			// Parse Button
			matches := buttonRegex.FindStringSubmatch(line)
			if len(matches) != 6 {
				return nil, prize, fmt.Errorf("failed to parse button: %s", line)
			}
			name := matches[1]
			x, _ := strconv.Atoi(matches[3])
			y, _ := strconv.Atoi(matches[5])
			buttons = append(buttons, Button{
				Name: name,
				Position: Position{
					X: x,
					Y: y,
				},
			})
		} else if strings.HasPrefix(line, "Prize") {
			// Parse Prize
			matches := prizeRegex.FindStringSubmatch(line)
			if len(matches) != 5 {
				return nil, prize, fmt.Errorf("failed to parse prize: %s", line)
			}
			x, _ := strconv.Atoi(matches[2])
			y, _ := strconv.Atoi(matches[4])
			prize = Prize{
				Position: Position{
					X: x,
					Y: y,
				},
			}
		}
	}

	return buttons, prize, nil
}
