package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	allArgs := strings.Join(args, ",")

	message := "Hello there!\n"

	fmt.Fprintf(os.Stderr, message)
	fmt.Fprintf(os.Stderr, allArgs)
}
