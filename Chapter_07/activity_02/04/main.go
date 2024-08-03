package main

import "fmt"

func main() {
	var str interface{} = "some thing"
	var i interface{} = 42
	var b interface{} = true
	fmt.Println(str)
	fmt.Println(i)
	fmt.Println(b)
}
