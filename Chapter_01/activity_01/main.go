package main

import (
	"fmt"
)

func main2() {
	helloList := []string{
		"Hello, world!",
	}

	fmt.Println(len(helloList))
	fmt.Println(helloList[len(helloList)-1])
	fmt.Println(helloList[len(helloList)])
}
