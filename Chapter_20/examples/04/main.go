package main

import "fmt"

// add returns the total of two integers added together
func Add(a, b int) int {
	return a + b
}

// multiply returns the total of one integer multiplied by the other
func Multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(Add(1, 1))
	fmt.Println(Multiply(2, 2))
}
