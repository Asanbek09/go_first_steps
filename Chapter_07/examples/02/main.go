package main

import "fmt"

type Speaker interface {
	Speak() string
}

type cat struct {
	name string
	age  int
}

func main() {
	c := cat{name: "Oreo", age: 9}
	fmt.Println(c.Speak())
	fmt.Println(c)
}

func (c cat) Speak() string {
	return "Purr Meow"
}

func (c cat) String() string {
	return fmt.Sprintf("%v (%v years old)", c.name, c.age)
}
