package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ch := make(chan string)
	signalStop := make(chan os.Signal, 1)
	urls := []string{
		"https://example.com",
		"https://google.com",
		"https://github.com",
		"https://ya.ru/?utm_referrer=https%3A%2F%2Fwww.google.com%2F",
	}

	signal.Notify(signalStop, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func() {
		for {
			<-signalStop
			signal.Stop(signalStop)
			log.Println("Signal type : ")
			//_, cancel := context.Canceled
			os.Exit(0)
		}
	}()

	for _, url := range urls {
		go crawlWeb(url, ch)
		fmt.Println(<-ch)
	}
}

func Signal(ch chan os.Signal) {
	for {
		<-ch
		signal.Stop(ch)
		//_, cancel := context.Canceled
		os.Exit(0)
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
