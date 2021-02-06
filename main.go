package main

import (
	"errors"
	"fmt"

	"net/http"
)

var errRequestFaild = errors.New("Request failed")

func main() {

	var results = make(map[string]string)

	urls := []string{

		"https://nomadcoders.co/",
		"https://www.naver.com/",
		"https://www.google.com/",
		"https://www.youtube.com/",
		"https://github.com/",
		"https://www.amazon.com/",
	}

	results["hello"] = "Hello"

	for _, url := range urls {

		result := "OK"
		err := hitURL(url)

		if err != nil {

			result = "FAILED"
		}

		results[url] = result

	}

	for url, result := range results {
		fmt.Println(url, result)
	}

}

func hitURL(url string) error {

	fmt.Println("Checking :", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {

		fmt.Println(err, resp.StatusCode)
		return errRequestFaild

	}

	return nil

}
