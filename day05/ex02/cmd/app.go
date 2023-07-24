package main

import (
	present "day05/ex02/struct"
	"fmt"
	"sort"
)

func getNCoolestPresents(pq present.PresentHeap, n int) present.PresentHeap {
	sort.Sort(pq)
	tmp := make(present.PresentHeap, n)
	for i := 0; i < n; i++ {
		tmp[i] = pq[i]
	}
	return tmp
}

func main() {
	h := present.PresentHeap{
		present.Present{Value: 5, Size: 1},
		present.Present{Value: 4, Size: 5},
		present.Present{Value: 3, Size: 1},
		present.Present{Value: 5, Size: 2},
	}
	fmt.Println(getNCoolestPresents(h, 1))
}
