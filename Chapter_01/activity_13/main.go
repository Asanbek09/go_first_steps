package main

import (
	"fmt"
	"time"
)

func main() {
	var count1 *int
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp
	t := &time.Time{}

	fmt.Printf("Count1: %#v \n", count1)
	fmt.Printf("Count2: %#v \n", count2)
	fmt.Printf("count3: %#v \n", count3)
	fmt.Printf("time : %#v \n", t)
}
