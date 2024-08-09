package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	onlyAfter, err := time.Parse(time.RFC3339, "2024-08-09T10:35:20+00:00")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(now, onlyAfter)
	fmt.Println(now.After(onlyAfter))
	if now.After(onlyAfter) {
		fmt.Println("Executing action!")
	} else {
		fmt.Println("Now is not the time yet!")
	}
}
