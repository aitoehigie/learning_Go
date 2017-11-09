package main

import (
	"fmt"
)

var (
	message = "The answer to life is %d\n"
	answer  = 42
	pi      = 3.14
)

func main() {
	fmt.Printf(message, answer, "while pi is", pi)
}
