package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var filename = "input.txt"

type Position struct {
	X int64
	Y int64
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

func solve() int64 {
	entries := readInput()

	var counter int64
	for _, entry := range entries {
		if entry.Prize.Position.X/entry.B.Position.X == entry.Prize.Position.Y/entry.B.Position.Y && entry.A.Position.Y/entry.B.Position.Y == entry.A.Position.X/entry.B.Position.X {
			fmt.Println("\n -------------------------- \n")
			fmt.Println(entry)
		}
		counter += calculate(entry)
	}

	// fmt.Println(counter)

	return counter
}

func calculate(entry Entry) int64 {
	// fmt.Println(entry)
	//
	term1 := new(big.Rat)
	term2 := new(big.Rat)
	term3 := new(big.Rat)
	term4 := new(big.Rat)
	rationalString := fmt.Sprintf("%d/%d", entry.Prize.Position.X, entry.A.Position.X)
	term1.SetString(rationalString)
	rationalString = fmt.Sprintf("%d/%d", entry.Prize.Position.Y, entry.A.Position.Y)
	term2.SetString(rationalString)
	rationalString = fmt.Sprintf("%d/%d", entry.B.Position.X, entry.A.Position.X)
	term3.SetString(rationalString)
	rationalString = fmt.Sprintf("%d/%d", entry.B.Position.Y, entry.A.Position.Y)
	term4.SetString(rationalString)
	b := new(big.Rat).Quo(new(big.Rat).Sub(term1, term2), new(big.Rat).Sub(term3, term4))
	if b.IsInt() {
		big1 := new(big.Int).SetInt64(entry.Prize.Position.X)
		bigMultiplier := new(big.Int).SetInt64(entry.B.Position.X)
		big3 := new(big.Int).SetInt64(entry.A.Position.X)

		temp := new(big.Rat).Mul(b, new(big.Rat).SetInt(bigMultiplier))

		// Ensure the denominator is 1 before converting to big.Int
		y := new(big.Int)
		if temp.IsInt() {
			// Convert directly to big.Int if temp is an integer
			y.Set(temp.Num())
		} else {
			// Handle cases where the denominator is not 1
			panic("The result is not an integer")
		}

		a := new(big.Int).Sub(big1, y) // x - y
		a.Div(a, big3)

		// fmt.Println(a, b.Num())
		A := a.Int64()
		B := b.Num().Int64()
		if A*entry.A.Position.X+B*entry.B.Position.X == entry.Prize.Position.X && A*entry.A.Position.Y+B*entry.B.Position.Y == entry.Prize.Position.Y {
			return 3*A + B
		} else {

			fmt.Println("WE GOT ONE!")
			fmt.Println(A, B)
		}
	}
	// b := 0

	//if isIntegral(b) {
	//	b = math.Round(b)
	//	fmt.Println("B: ", int64(b))
	//	a = (entry.Prize.Position.Y - b*entry.B.Position.Y) / entry.A.Position.Y
	//	// fmt.Println(a, b)
	//	return 3*int(a) + int(b)
	//}

	return 0
}

//func isIntegral(val int64) bool {
//	fmt.Println("B: ", val)
//	return math.Abs(val-math.Round(val)) < floating
//}

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
					X: int64(x),
					Y: int64(y),
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
			// prize = Prize{
			// 	Position: Position{
			// 		X: int64(x),
			// 		Y: int64(y),
			// 	},
			// }
			prize = Prize{
				Position: Position{
					X: int64(x + 10000000000000),
					Y: int64(y + 10000000000000),
				},
			}
		}
	}

	return buttons, prize, nil
}
