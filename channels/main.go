package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range urls {

		go checkLink(link, c)

	}
	for l := range c {
		go func(lin string) {
			time.Sleep(5 * time.Second)
			checkLink(lin, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {

		fmt.Println(link, "might be down. The Error is : ", err)
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}
