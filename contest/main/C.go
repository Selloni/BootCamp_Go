package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	write := bufio.NewWriter(os.Stdout)
	defer write.Flush()

	var (
		countLine int
		countNum  int
		//buff      []int
		lvl int
	)
	fmt.Fscan(read, &countLine)
	for i := 0; i < countLine; i++ {
		var tmpBuff []int
		list := make(map[int]int)
		fmt.Fscan(read, &countNum)
		for j := 1; j < countNum+1; j++ {
			fmt.Fscan(read, &lvl)
			list[j] = lvl
		}
		for key := range list {
			tmpBuff = append(tmpBuff, key)
		}

		sort.Ints(tmpBuff)
		fmt.Println(tmpBuff)
		ll := len(list)
		for m := 1; m < ll; m++ {
			fmt.Fprintln(write, findKey(&list, m), findKey(&list, m+1))
		}
	}
}

func findKey(list *map[int]int, val int) int {
	for key, value := range *list {
		if value == val || value == val+1 || value == val-1 {
			delete(*list, key)
			return key
		}
	}
	return 0
}
