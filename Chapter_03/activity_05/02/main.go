package main

import "fmt"

func main() {
	comment1 := `This is the best thing ever!`
	comment2 := `This is the Best \n thing ever!`
	comment3 := "This is the Best \n thing ever!"
	fmt.Println(comment1, "\n\n")
	fmt.Println(comment2, "\n\n")
	fmt.Println(comment3, "\n")
}
