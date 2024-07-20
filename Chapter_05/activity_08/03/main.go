package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Executing an anonymous function using a variable")
	}
	fmt.Println("Line after anonymous function declaration")
	f()
}
