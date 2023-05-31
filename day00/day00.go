package main

import (
	"fmt"
	"log"
	"math"
	"sort"
)

func main() {
	var sum, num float64
	var buff []float64
	for {
		_, err := fmt.Scanf("%f", &num)
		if err != nil {
			break
		}
		if num > -100000 && num < 100000 {
			buff = append(buff, num)
			sum += num
		} else {
			log.Fatal("range of numbers -100000, +100000")
		}
	}
	if len(buff) != 0 {
		fmt.Printf("Mean: %.2f\n", mean(sum, float64(len(buff))))
		if (len(buff) % 2) != 0 {
			m, _ := median(buff)
			fmt.Printf("Median: %.2f\n", m)
		} else {
			m, me := median(buff)
			fmt.Printf("Median: %.2f %.2f\n", m, me)
		}
		fmt.Printf("Mode %.2f\n", mode(buff))
		fmt.Printf("SD: %.2f\n", standardDeviation(buff, sum))
	} else {
		log.Fatal("incorrect input")
	}
}

func mean(sum float64, tmp float64) float64 {
	return sum / tmp
}

func median(buff []float64) (float64, float64) {
	sort.Float64s(buff)
	if (len(buff) % 2) != 0 {
		return buff[len(buff)/2], buff[len(buff)/2]
	} else {
		round := math.Round(float64(len(buff) / 2))
		return buff[int(round-1)], buff[int(round)]
	}
}

func mode(buff []float64) float64 {
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
		} else if value == count {
			if max > key {
				max = key
			}
		}
	}
	return max
}

func standardDeviation(buff []float64, sum float64) float64 {
	var itog float64
	mean := sum / float64(len(buff))
	for _, exp := range buff {
		itog += math.Pow(exp-mean, 2)
	}
	itog = math.Sqrt(itog) / float64(len(buff))

	return itog
}
