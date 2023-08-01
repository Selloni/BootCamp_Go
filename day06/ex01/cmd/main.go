package main

import "fmt"

type Post struct {
	Id   int
	Text string
}

func main() {
	fmt.Println("start:8888port")
	handleRequest()
}
