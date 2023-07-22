package main

import (
	"day05/ex00/tree"
	"fmt"
)

func main() {
	node := tree.TreeNode{}
	node.FillNode(true, true)
	value := tree.AreToysBalanced(&node)
	fmt.Println(value)
}
