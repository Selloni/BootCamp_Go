package main

import (
	"day07/app"
	"fmt"
	"sort"
)

func main() {
	tmp := []int{1, 5, 10}
	sort.Ints(tmp)
	fmt.Println(app.MinCoins(13, tmp))
}
