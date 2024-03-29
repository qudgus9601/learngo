package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://www.academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}
	for i:=0; i< len(urls);i++{
		result := <- c
		results[result.url] = result.status
	}
	
	for url,result := range results {
		fmt.Println(url, result);
	}
}

func hitURL(url string, c chan requestResult) {
	response, err := http.Get(url)
	status := "OK"
	if err != nil || response.StatusCode >= 400 {
		status = "FAILED"
	}

	c <- requestResult{url:url,status: status}
}