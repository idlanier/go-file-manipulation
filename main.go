package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	myFile, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Printf("Open File Error : %v", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, myFile)

}
