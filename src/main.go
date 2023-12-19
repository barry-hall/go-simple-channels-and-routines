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

	for _, link := range links {
		checkLink(link)
	}

	duration := time.Since(start)

	roundedDuration := Round(duration.Seconds(), 2)

	fmt.Printf("Time taken: %.2f seconds\n", roundedDuration)
}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}

func Round(num float64, decimalPlaces int) float64 {
	pow := math.Pow10(decimalPlaces)
	return math.Round(num*pow) / pow
}
