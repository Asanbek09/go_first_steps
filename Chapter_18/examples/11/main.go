package main

import "fmt"

func readThem(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	ch := make(chan int)
	go readThem(ch)
	ch <- 1
	ch <- 2
	ch <- 3
}
