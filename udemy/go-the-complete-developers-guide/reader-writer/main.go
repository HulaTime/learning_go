package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, httpErr := http.Get("https://www.google.com")
	if httpErr != nil {
		fmt.Println("Err:", httpErr)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)
}
