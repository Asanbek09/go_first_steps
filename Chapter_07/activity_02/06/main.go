package main

import "fmt"

func main() {
	var str interface{} = "the book club"
	v, isValid := str.(int)
	fmt.Println(v, isValid)
}
