package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "found me"
	if strings.Contains(str, "found") {
		fmt.Println("value found in str")
	}
}
