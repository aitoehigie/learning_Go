package main

import (
	"fmt"
)

var (
	atoz = "the quick brown fox jumps over the lazy dog\n"
)

func main() {
	fmt.Printf("Slicing the string: %s\n", atoz[15:])
}
