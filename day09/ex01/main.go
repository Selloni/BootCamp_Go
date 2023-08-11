package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	var ch chan string

	urls := []string{
		"https://example.com",
		"https://google.com",
		"https://github.com",
		"https://ya.ru/?utm_referrer=https%3A%2F%2Fwww.google.com%2F",
	}

	for _, url := range urls {
		resp, err := crawlWeb(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp)
	}
}

func crawlWeb(url string) (string, error) {
	var client http.Client
	var bodyString string

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString = string(bodyBytes)
	}
	return bodyString, nil
}
