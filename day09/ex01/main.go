package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := crawlWeb("https://journal.tinkoff.ru/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
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
