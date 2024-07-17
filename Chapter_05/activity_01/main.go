package main

import "fmt"

func checkNumber() {
	for i := 1; i < 30; i++ {
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}

func main() {
	fmt.Println("Start")
	checkNumber()
	fmt.Println("End")
}
