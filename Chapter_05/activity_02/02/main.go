package main

import "fmt"

func main() {
	checkNumbers(10)
}

func checkNumbers(end int) {
	for i := 1; i <= end; i++ {
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
