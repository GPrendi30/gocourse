package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://facebook.com",
		"http://google.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		// function
		go checkLink(link, c)

	}

	for link := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(link)
	}
}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link + " might be down i think"
		return
	}

	fmt.Println(link, "is up!")
	c <- link

}
