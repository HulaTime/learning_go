package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, l := range links {
		go checkLink(l, c)
	}

	for l := range c {
		func(link string) {
			time.Sleep(5 * time.Second)
			go checkLink(link, c)
		}(l)
	}
}

func checkLink(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down", err)
		c <- url
	}
	fmt.Println(url, "Link is up")
	c <- url
}
