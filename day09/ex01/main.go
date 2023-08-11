package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("keep up to 8 urls")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	resp, err := crawlWeb(scanner.Text())
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
