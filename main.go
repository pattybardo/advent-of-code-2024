package main

import (
	"fmt"
	"strings"
)

func main() {
	x := "000100"

	fmt.Println(strings.TrimLeft(x, "0"))
}
