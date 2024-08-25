package main

import "fmt"

func main() {
	helloString := "Hello"
	packtString := "Packt"
	joinString := fmt.Sprintf("%s %s", helloString, packtString)
	fmt.Println(joinString)
}
