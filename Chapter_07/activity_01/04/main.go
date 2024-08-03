package main

import "fmt"

type Speaker interface {
	Speak() string
}

type cat struct{}

func main() {
	c := cat{}
	catSpeak(c)
}

func (c cat) Speak() string {
	return "Purr Meow 3"
}

func catSpeak(c cat) {
	fmt.Println(c.Speak())
}
