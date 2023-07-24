package main

import (
	"fmt"
	"sort"
)

type PresentHeap []Present

type Present struct {
	Value int
	Size  int
}

func (pq PresentHeap) Len() int {
	return len(pq)
}

func (pq PresentHeap) Less(i, j int) bool {
	if pq[i].Value > pq[j].Value {
		return true
	}
	if pq[i].Value == pq[j].Value && pq[i].Size < pq[i].Size {
		return true
	}
	return false
}

func (pq PresentHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func getNCoolestPresents(pq PresentHeap, n int) PresentHeap {
	sort.Sort(pq)
	tmp := make(PresentHeap, n)
	for i := 0; i < n; i++ {
		tmp[i] = pq[i]
	}
	return tmp
}

func main() {
	h := PresentHeap{
		Present{Value: 5, Size: 1},
		Present{Value: 4, Size: 5},
		Present{Value: 3, Size: 1},
		Present{Value: 5, Size: 2},
	}
	fmt.Println(getNCoolestPresents(h, 1))
}
