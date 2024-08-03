package main

import "fmt"

type Speaker interface {
	Speak() string
}

type cat struct{}

func main() {
	c := cat{}
	chatter(c)
}

func (c cat) Speak() string {
	return "Purr Meow 2"
}

func chatter(s Speaker) {
	fmt.Println(s.Speak())
}
