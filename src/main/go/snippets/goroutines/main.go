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
		"http://golang.com",
		"http://reddit.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go makeReques(url, c)
	}

	// Alternative syntax
	// for {
	// 	go makeReques(<-c, c)
	// }

	// for l := range c {
	// 	go makeReques(l, c)
	// }

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			makeReques(link, c)
		}(l)
	}
}

func makeReques(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		c <- url
		return
	}

	fmt.Printf("Url %v is up\n", url)
	c <- url
}
