package main

import (
	"fmt"
	"time"
)

func main() {
	buff := []int{1, 1, 3, 7, 4, 5}
	fmt.Print(sleepSort(buff))
}

func sleepSort(x []int) (sorti []int) {
	ch := make(chan int)

	for _, i := range x {
		go forSleep(i, ch)
	}

	for i := 0; i < len(x); i++ {
		sorti = append(sorti, <-ch)
	}
	return
}

func forSleep(x int, ch chan int) {
	time.Sleep(time.Duration(x) * time.Second)
	ch <- x
}
