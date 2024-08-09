package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	losAngeles, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The local current times is: ", current.Format(time.ANSIC))
	fmt.Println("The time in Los Angeles is: ", current.In(losAngeles).Format(time.ANSIC))
}
