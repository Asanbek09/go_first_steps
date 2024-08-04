package main

import (
	"fmt"

	"github.com/sicoyle/printer/printer"
)

func main() {
	msg := printer.PrintNewUUID()
	fmt.Println(msg)
}
