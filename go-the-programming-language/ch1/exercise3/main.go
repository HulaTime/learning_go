package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var args []string

	for i := 1; i <= 100_000_000; i++ {
		args = append(args, fmt.Sprintf(" %d", i))
	}

	startTime := time.Now()
	fmt.Println(fmt.Sprintf("Start time is %d", startTime.Unix()))

	s, sep := "", ""
	for _, val := range os.Args[1:] {
		s += sep + val
		sep = " "
	}
	endTime := time.Now()
	fmt.Println(fmt.Sprintf("endTime time is %d", endTime.Unix()))

	fmt.Println("Time spent %s", endTime.Unix()-startTime.Unix())

	fmt.Println(fmt.Sprintf("Start time is %d", startTime.Unix()))
	strings.Join(args, "")
	fmt.Println(fmt.Sprintf("Start time is %d", endTime.Unix()))
	fmt.Println("Time spent %s", endTime.Unix()-startTime.Unix())
}
