package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
	fmt.Println("The filename to read is: ", filename)
	file, _ := os.Open(filename)
	io.Copy(os.Stdout, file)
}
