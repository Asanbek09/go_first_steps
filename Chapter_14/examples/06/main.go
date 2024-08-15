package main

import (
	"fmt"
	"os"
)

func main() {
	message := []byte("look!")
	err := os.WriteFile("test.txt", message, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
