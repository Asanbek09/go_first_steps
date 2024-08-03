package main

import "fmt"

func findMaxInt(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	max := findMaxInt([]int{1, 32, 5, 8, 10, 11})
	fmt.Printf("max integer value : %v\n", max)
	maxGenericInt := findMaxGeneric([]int{1, 32, 5, 8, 10, 11, 26})
	fmt.Printf("max generic value : %v\n", maxGenericInt)
}

func findMaxFloat(nums []float64) float64 {
	if len(nums) == 0 {
		return -1
	}

	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func findMaxGeneric[Num int | float64](nums []Num) Num {
	if len(nums) == 0 {
		return -1
	}

	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
