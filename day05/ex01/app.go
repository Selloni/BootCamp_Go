package main

import (
	tree "day05/ex00/tree"
	"fmt"
)

func unrollGarland(node tree.TreeNode) {
	buff := make(map[int][]tree.TmpType)

	fl := tree.Height(&node)
	i := int(fl)
	if i%2 == 0 {
		i++
	}
	takeTree(&node, &buff, i)
	var sl []tree.TmpType
	for ; i > 0; i-- {
		sl = append(sl, buff[i]...)
	}
	fmt.Print(sl)
}

func takeTree(node *tree.TreeNode, buff *map[int][]tree.TmpType, i int) {
	if node == nil {
		return
	}
	takeTree(node.Left, buff, i-1)
	takeTree(node.Right, buff, i-1)

	if i%2 != 0 {
		(*buff)[i] = reverse((*buff)[i])
	}
	(*buff)[i] = append((*buff)[i], node.HasToy)
}

func reverse(sl []tree.TmpType) []tree.TmpType {
	var buff []tree.TmpType
	for l := len(sl) - 1; l > -1; l-- {
		buff = append(buff, sl[l])
	}
	return buff
}

func main() {
	node := tree.TreeNode{
		HasToy: true,
	}
	node.FillNode(true, true, true, false, false, true)
	//node.FillNode(2, 3, 4, 5, 6, 7)
	unrollGarland(node)

}
