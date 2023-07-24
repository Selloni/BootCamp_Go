package main

import (
	present "day05/ex02/struct"
)

func main() {
	h := present.PresentHeap{
		present.Present{Value: 5, Size: 1},
		present.Present{Value: 4, Size: 5},
		present.Present{Value: 3, Size: 1},
		present.Present{Value: 5, Size: 2},
	}

}

func grabPresents(item present.PresentHeap, size int) {
	matrix := make([][]int, item.Len()+1)
	for i := range matrix {
		matrix[i] = make([]int, size+1)
	}
}
