package main

import "fmt"

func sum(from, to int) int {
	res := 0
	for i := from; i <= to; i++ {
		res += i
	}
	return res
}

func main() {
	s1 := sum(1, 100)
	fmt.Println(s1)
}
