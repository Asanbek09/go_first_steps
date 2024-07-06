package main

import "fmt"

var level = "pkg"

func main() {
	fmt.Println("main start: ", level)
	if true {
		fmt.Println("Block start: ", level)
		funcA()
	}
}

func funcA() {
	fmt.Println("funcA start: ", level)
}
