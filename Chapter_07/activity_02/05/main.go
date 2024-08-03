package main

import (
	"fmt"
	"strings"
)

func main() {
	var str interface{} = "some string"
	v := str.(string)
	fmt.Println(strings.Title(v))
}
