package main

import "fmt"

func main() {
	logLevel := "круто что так просто"
	for index, runeVal := range logLevel {
		fmt.Println(index, string(runeVal))
	}
}
