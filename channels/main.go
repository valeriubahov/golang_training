package main

import (
	"fmt"
	"net/http"
	"time"
)

//time package in this case is used to pause the execution of the code for a selected period of time

func main() {
	links := []string{
		"http://facebook.com",
		"http://youtube.com",
		"http://amazon.com",
		"http://google.com",
		"http://linkedin.com",
	}

	// use the make function to create a channel that will transport only values of type string
	c := make(chan string)

	for _, v := range links {
		// the `go` operator will create a new goroutine which are light-weight processes that runs separatelly
		// we pass the channel as a parameter in order to transport the values from 1 proces to another
		go checkLink(v, c)
	}

	// this for loop is quiet strange
	// we use the range of c to listen to a channel and every time a value is transported with that channel we perform the loop
	for l := range c {
		// this is an IIF - Immediate Invoked Function
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// set the value into the channek
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
