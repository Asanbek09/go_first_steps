package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <name>")
		return
	}

	name := args[1]
	greetings := fmt.Sprintf("Hello, %s! Welcome to the command line", name)
	fmt.Println(greetings)
}
