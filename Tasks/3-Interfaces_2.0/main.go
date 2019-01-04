package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("ERROR: You must put a file name")
		os.Exit(1)
	}
	fileName := args[1]
	fmt.Println("File Name:", fileName)
	file, error := os.Open(fileName)
	if error != nil {
		fmt.Println("Error reading the file:", error)
		os.Exit(1)
	}
	fmt.Println("File Content:")
	io.Copy(os.Stdout, file)
}
