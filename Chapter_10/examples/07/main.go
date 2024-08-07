package main

import "fmt"

var budgetCategories = make(map[int]string)

func init() {
	fmt.Println("Initizlizing our budgetCategories")
	budgetCategories[1] = "Car insurance"
	budgetCategories[2] = "Mortgage"
	budgetCategories[3] = "Electricity"
	budgetCategories[4] = "Retirement"
	budgetCategories[5] = "Vacation"
	budgetCategories[7] = "Groceries"
	budgetCategories[8] = "Car Payment"
}

func main() {
	for k, v := range budgetCategories {
		fmt.Printf("key: %d, value: %s\n", k, v)
	}
}
