package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	num int
	lvl int
}

type SortPair []pair

func (a SortPair) Len() int           { return len(a) }
func (a SortPair) Less(i, j int) bool { return a[i].lvl < a[j].lvl }
func (a SortPair) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	read := bufio.NewReader(os.Stdin)
	write := bufio.NewWriter(os.Stdout)
	defer write.Flush()

	var (
		countLine int
		countNum  int
		lvl       int
		//tmp       pair
		pp []pair
	)
	fmt.Fscan(read, &countLine)
	for i := 0; i < countLine; i++ {
		fmt.Fscan(read, &countNum)
		for j := 1; j < countNum+1; j++ {

		}
	}
}
