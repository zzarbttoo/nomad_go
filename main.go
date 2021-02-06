package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFaild = errors.New("Request failed")

type requestResult struct {
	url    string
	status string
}

func main() {

	results := make(map[string]string)
	channel := make(chan requestResult)

	urls := []string{

		"https://nomadcoders.co/",
		"https://www.naver.com/",
		"https://www.google.com/",
		"https://www.youtube.com/",
		"https://github.com/",
		"https://www.amazon.com/",
	}

	for _, url := range urls {

		go hitURL(url, channel)
	}

	//receiving message

	for i := 0; i < len(urls); i++ {
		result := <-channel
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

//write direction
//can only write
func hitURL(url string, channel chan<- requestResult) {

	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {

		status = "FAILED"

	}

	channel <- requestResult{url: url, status: status}

	// can send to the channel c<-result{}
	// but can also recieve from channel

}
