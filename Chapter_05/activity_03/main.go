package main

import "fmt"

func main() {
	greeting("Cayden", 45)
}

func greeting(name string, age int) {
	fmt.Printf("%s is %d", name, age)
}
