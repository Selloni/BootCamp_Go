package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	write := bufio.NewWriter(os.Stdout)
	some := 0
	fmt.Fscan(read, &some)
	//for i := 0; i < some; i++ {

	//}
	defer write.Flush()
}
