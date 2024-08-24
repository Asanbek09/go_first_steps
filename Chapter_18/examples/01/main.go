package main

import "fmt"

func hello() {
	fmt.Println("Hello World")
}

func main() {
	fmt.Println("start")
	go hello()
	fmt.Println("end")
}
