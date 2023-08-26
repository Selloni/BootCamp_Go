package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	write := bufio.NewWriter(os.Stdout)
	defer write.Flush()

	var (
		countLine int
		countNum  int
		//buff      []int
		tmpBuff []int
		lvl     int
	)
	fmt.Fscan(read, &countLine)
	for i := 0; i < countLine; i++ {
		list := make(map[int]int)
		fmt.Fscan(read, &countNum)
		for j := 0; j < countNum; j++ {
			fmt.Fscan(read, &lvl)
			list[j] = lvl
		}

		//for i, j := range buff {
		//	tmp =
		//	//tmpBuff = append(tmpBuff, i)
		//}
		fmt.Fprintln(write, tmpBuff)
	}
}

func findPair(buff int, tmp int, write *bufio.Writer) {
	if tmp == buff {
		fmt.Fprintln(write, tmp, buff)
	}

}
