package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	ch := make(chan string)
	urls := []string{
		"https://example.com",
		"https://google.com",
		"https://github.com",
		"https://ya.ru/?utm_referrer=https%3A%2F%2Fwww.google.com%2F",
	}
	for _, url := range urls {
		go crawlWeb(url, ch)
		fmt.Println(<-ch)
	}
}

func crawlWeb(url string, ch chan string) {
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		ch <- string(bodyBytes)
	} else {
		<-ch
	}
}
