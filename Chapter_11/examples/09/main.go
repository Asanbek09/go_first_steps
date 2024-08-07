package main

import "fmt"

type person struct {
	lname  string
	age    int
	salary float64
}

func main() {
	fname := "Joe"
	grades := []int{100, 87, 74}
	states := map[string]string{"KY": "Kentucky", "WV": "West Virginia", "VA": "Virginia"}
	p := person{lname: "Lincoln", age: 210, salary: 25000.0}
	fmt.Printf("fname value is %v\n", fname)
	fmt.Printf("grades value is %v\n", grades)
	fmt.Printf("states value is %v\n", states)
	fmt.Printf("p value is %v\n", p)
}
