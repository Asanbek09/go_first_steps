package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "", "File name")
	flag.Parse()
	file, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s: File does not exist\n", name)
			fmt.Println(file)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("file name: %s\n IsDir: %t\n ModTime: %v\n Mode: %v\n Size: %d\n", file.Name(), file.IsDir(), file.ModTime(), file.Mode(), file.Size())
}
