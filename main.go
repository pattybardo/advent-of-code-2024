package main

import "fmt"

func main() {
	var test []int
	for range 2 {
		test[1] = 2
		fmt.Println("Hello, Go!"[1:])
		fmt.Println(test)
	}
}
