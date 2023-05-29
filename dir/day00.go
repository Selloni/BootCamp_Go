package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var sum, num float64
	var buff []float64
	for {
		_, err := fmt.Scanln(&num)
		if err != nil {
			break
		}
		buff = append(buff, num)
		sum += num
	}
	mean(sum, float64(len(buff)))
	median(buff)
	mode(buff)
	standardDeviation(buff, sum)
}

func mean(sum float64, tmp float64) {
	fmt.Printf("Mean: %.2f\n", sum/tmp)
}

func median(buff []float64) {
	sort.Float64s(buff)
	if (len(buff) % 2) != 0 {
		fmt.Printf("Median: %.2f\n", buff[len(buff)/2])
	} else {
		round := math.Round(float64(len(buff) / 2))
		fmt.Printf("Median: %.2f %.2f\n", buff[int(round-1)], buff[int(round)])
	}
}

func mode(buff []float64) {
	countMap := make(map[float64]int)
	for _, value := range buff {
		countMap[value] += 1
	}
	max := buff[0]
	count := 0
	for _, key := range buff {
		value := countMap[key]
		if value > count {
			count = value
			max = key
		}
	}
	fmt.Printf("Mode %.2f\n", max)
}

func standardDeviation(buff []float64, sum float64) {
	var itog float64
	mean := sum / float64(len(buff))
	for _, exp := range buff {
		itog += math.Pow(2, (exp - mean))
	}
	itog = math.Sqrt(itog) / float64(len(buff))
	fmt.Printf("SD: %.2f\n", itog)
}
