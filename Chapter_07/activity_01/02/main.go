package main

import "fmt"

type Speaker interface {
	Speak() string
}
type cat struct{}

func main() {
	c := cat{}
	fmt.Println(c.Speak())
}

func (c cat) Speak() string {
	return "Purr Meow"
}
