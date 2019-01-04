package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string) //Create a channel of type string

	for _, link := range links {
		go checkLink(link, c) //We pass the channel to the function
	}
	/*for {
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(<-c, c)
		}()
	}*/
	for l := range c { //waiting for the channel callback and assign the value to l (link)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		log.Println(link, "might be down!")
		c <- link
		return
	}
	log.Println(link, "is up!")
	c <- link
}
