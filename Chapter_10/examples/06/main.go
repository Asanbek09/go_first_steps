package main

import "fmt"

var name = "Gopher"

func init() {
	fmt.Println("Hello, ", name)
}

func main() {
	fmt.Println("Hello, main function")
}
