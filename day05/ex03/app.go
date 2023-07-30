package main

import (
	present "day05/ex02/struct"
	"fmt"
	"math"
)

func main() {
	h := present.PresentHeap{
		present.Present{Value: 1, Size: 1},
		present.Present{Value: 4, Size: 5},
		present.Present{Value: 3, Size: 3},
		present.Present{Value: 2, Size: 1},
	}
	fmt.Println(grabPresents(h, 1))
}

func grabPresents(item present.PresentHeap, disk int) int {
	var (
		value []int
		size  []int
	)
	for _, press := range item {
		value = append(value, press.Value)
		size = append(size, press.Size)
	}

	matrix := make([][]int, item.Len()+1)
	for i := range matrix {
		matrix[i] = make([]int, disk+1)
	}
	for i := 0; i <= item.Len(); i++ {
		for j := 0; j <= disk; j++ {
			if i == 0 || j == 0 {
				matrix[i][j] = 0
			} else {
				if size[i-1] > j {
					matrix[i][j] = matrix[i-1][j]
				} else {
					prev := matrix[i-1][j]
					tmp := value[i-1] + matrix[i-1][j-size[i-1]]
					maxValue := math.Max(float64(prev), float64(tmp))
					matrix[i][j] = int(maxValue)
				}
			}
		}
	}
	return matrix[item.Len()][disk]
}
