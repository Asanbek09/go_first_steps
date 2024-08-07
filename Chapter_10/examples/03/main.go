package main

import (
	"fmt"
)

func main() {
	str := "found me"
	slc := strings.explode(str, 3)
	fmt.Println(slc)
}
