package main

import (
	"fmt"
	"math"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	start := time.Now()

	// create a channel
	c := make(chan string)

	for _, link := range links {
		// pass the channel
		go checkLink(link, c)
	}

	for l := range c {
		// recieve from the channel
		go checkLink(l, c)
	}

	duration := time.Since(start)
	roundedDuration := Round(duration.Seconds(), 2)
	fmt.Printf("Time taken: %.2f seconds\n", roundedDuration)
}

// channel is passed as an argument and includes a type
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// send to the channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// send to the channel
	c <- link
}

func Round(num float64, decimalPlaces int) float64 {
	pow := math.Pow10(decimalPlaces)
	return math.Round(num*pow) / pow
}
