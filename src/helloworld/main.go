package main

import (
	"fmt"
)

var (
	atoz          = "the quick brown fox jumps over the lazy dog\n"
	len_of_string = len(atoz)
)

func main() {
	for i, c := range atoz {
		fmt.Printf("%d %c\n", i, c)
	}
	fmt.Printf("The string is of length: %d\n", len_of_string)
}
