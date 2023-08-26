package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	write := bufio.NewWriter(os.Stdout)
	defer write.Flush()
	var (
		countLine int
		countNum  int
		value     string
	)
	fmt.Fscan(read, &countLine)
	for i := 0; i < countLine; i++ {
		fmt.Fscan(read, &countNum)
		goods := make(map[string]int)
		for j := 0; j < countNum; j++ {
			fmt.Fscan(read, &value)
			goods[value] = goods[value] + 1
		}
		var sum int
		for v, count := range goods {
			vint, _ := strconv.Atoi(v)
			for i := count; i != 0; i-- {
				if i%3 == 0 {
					continue
				}
				sum = sum + vint
			}
		}
		fmt.Fprintln(write, sum)
	}
}

//func main() {
//	read := bufio.NewReader(os.Stdin)
//	write := bufio.NewWriter(os.Stdout)
//	defer write.Flush()
//	input, err := read.ReadString('\n')
//	if err != nil {
//		log.Fatal(err)
//	}
//	//some := 0
//	//fmt.Fscan(read, &some)
//	goods := make(map[string]int)
//	sl := strings.Fields(input)
//	for _, i := range sl {
//		goods[i] = goods[i] + 1
//	}
//	var sum int
//	for v, count := range goods {
//		vint, _ := strconv.Atoi(v)
//		for i := count; i != 0; i-- {
//			if i%3 == 0 {
//				continue
//			}
//			sum = sum + vint
//
//		}
//	}
//	sumStr := strconv.Itoa(sum)
//	write.WriteString(sumStr)
//}
