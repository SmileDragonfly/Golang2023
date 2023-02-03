package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	c := make(chan int)
	for _, link := range links {
		go checkLink(link, c)
	}
	for _, _ = range links {
		<-c
	}
}

func checkLink(link string, c chan int) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- 1
		return
	}
	fmt.Println(link, "is up!")
	c <- 1
}
