package main

import (
	"fmt"
	"sync"
)

func main() {
	one := generator(17)
	two := generator(28)
	three := generator(13)
	out := multiplex(one, two, three)
	for i := range out {
		fmt.Printf("%d ", i)
	}
}

func multiplex(ch ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	out := make(chan interface{})

	send := func(c <-chan interface{}) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(ch))
	for _, c := range ch {
		go send(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func generator(data any) <-chan any {
	ch := make(chan any)
	go func() {
		for i := 0; i < 4; i++ {
			ch <- data
		}
		close(ch)
	}()
	return ch
}
