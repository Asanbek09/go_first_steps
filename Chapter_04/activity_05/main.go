package main

import "fmt"

func message() string {
	arr := [4]string{"ready", "get", "go", "to"}
	arr[1] = "It's"
	arr[0] = "time"
	return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
}

func main() {
	fmt.Println(message())
}
