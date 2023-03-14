package urlcheck

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {

	var results = map[string]string{}
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
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "Failed"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

var errRequestFailed error = errors.New("request failed")

func hitURL(url string) error {
	fmt.Println(url)
	response, err := http.Get(url)

	if err != nil || response.StatusCode >= 400 {
		fmt.Println(err)
		return errRequestFailed
	}

	return nil
}