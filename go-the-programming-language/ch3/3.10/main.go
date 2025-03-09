package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer

	for i, c := range s {
		if (i+1)%3 == 0 {
			buf.WriteString(fmt.Sprintf("%c,", c))
		} else {
			buf.WriteString(string(c))
		}
	}

	return buf.String()
}

func main() {
	fmt.Println(comma("1234567890"))
}
