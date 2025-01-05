package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, httpErr := http.Get("https://www.google.com")
	if httpErr != nil {
		fmt.Println("Err:", httpErr)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
