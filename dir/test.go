package main

import (
	"fmt"
	"sort"
	// "math"
)

func main() {
    var sum, num, tmp float64
    var buff[] float64
    for {
        count, err := fmt.Scanln(&num)
        if (err != nil){
            fmt.Println("stop")
            break    
        }
        count++
        buff = append(buff, num)
        sum += num
        tmp = float64(count)
        fmt.Printf("Mean: %f\n", mean(sum, tmp))
    }
    sum = sum/tmp
    fmt.Printf("Mean: %f\n", mean(sum, tmp))
    median(buff)
    fmt.Println(buff)
}

func mean (sum float64, tmp float64) float64 {
    return sum/tmp
}

func median (buff[] float64) {
    sort.Float64s(buff)
    if 
}