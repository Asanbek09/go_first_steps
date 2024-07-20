package main

import "fmt"

func main() {
	nums("James", 99, 100)
}

func nums(str string, i ...int) {
	fmt.Println(str)
	fmt.Println(i)
}
