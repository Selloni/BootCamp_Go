package main

import "sync"

func main() {
	out := multiplex()
}

func multiplex(ch ...chan interface{}) chan interface{} {
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
