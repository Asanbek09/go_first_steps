package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	buf := make([]byte, 1)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			fmt.Print(string(buf[:n]))
		}
	}
}
