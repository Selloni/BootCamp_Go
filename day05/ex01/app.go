package main

import (
	tree "day05/ex00/tree"
	"fmt"
)

func unrollGarland(node tree.TreeNode) {

}

func printTree(node *tree.TreeNode, buff *[]bool) {
	//i := 0
	//fmt.Print(node.HasToy)
	if node == nil {
		return
	}
	printTree(node.Left, buff)
	printTree(node.Right, buff)
	//if i%2 == 0 {
	//
	//} else {
	//
	//}
	*buff = append(*buff, node.HasToy)
}

//func returnBool(node tree.TreeNode) bool  {
//
//}

func reverse(sl []bool) []bool {
	var buff []bool
	for l := len(sl) - 1; l > -1; l-- {
		buff = append(buff, sl[l])
	}
	return buff
}

func main() {
	node := tree.TreeNode{
		HasToy: false,
	}
	node.FillNode(true, true, true, false, false)
	var buff []bool
	printTree(&node, &buff)
	fmt.Print(buff)
}
