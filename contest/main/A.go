package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	write := bufio.NewWriter(os.Stdout)

	var l, r int
	fmt.Fscan(read, &l, &r)
	fmt.Fprintln(write, sum(l, r))

	defer write.Flush()
}
func sum(a int, b int) int {
	return a + b
}
