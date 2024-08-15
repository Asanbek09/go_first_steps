package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File contents: ")
	fmt.Println(string(content))
}
