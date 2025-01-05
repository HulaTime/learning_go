package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("Only one file plz")
		os.Exit(1)
	}
	fb, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Err:", err)
		os.Exit(1)
	}

	io.Copy(os.Stderr, fb)
}
