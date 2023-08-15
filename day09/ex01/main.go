package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
)

func main() {
	forProf()
	//http.ListenAndServe("localhost:8080", nil)
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
		//for {
		<-signalStop
		//signal.Stop(signalStop)
		_, cancel := context.WithCancel(context.Background())
		//_, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		fmt.Printf("\nThe process has closed\n")
		//_, cancel := context.Canceled
		os.Exit(0)
		//}
	}()

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

func forProf() {
	file, err := os.Create("heap.profile")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	pprof.Lookup("heap").WriteTo(file, 0)
}
