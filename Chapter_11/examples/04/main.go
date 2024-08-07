package main

import "fmt"

func main() {
	fname := "Joe"
	gpa := 3.75
	hasJob := true
	age := 24
	hourlyWage := 45.54
	fmt.Printf("%s has a gpa of %f\n", fname, gpa)
	fmt.Printf("He has a job equals %t.\n", hasJob)
	fmt.Printf("he is %d earning %v per hour\n", age, hourlyWage)
}
