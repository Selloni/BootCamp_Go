package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	ch = make(chan string)

	urls := []string{
		"https://example.com",
		"https://google.com",
		"https://github.com",
		"https://ya.ru/?utm_referrer=https%3A%2F%2Fwww.google.com%2F",
	}

	for _, url := range urls {
		ch <- url
		err := crawlWeb(ch)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(<-ch)
	}
}

func crawlWeb(chUrl chan string) error {
	var client http.Client
	//var bodyString string
	url := <-chUrl
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		chUrl <- string(bodyBytes)
	} else {
		<-chUrl
	}
	return nil
}
