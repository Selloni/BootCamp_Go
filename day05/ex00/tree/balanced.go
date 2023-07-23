package tree

import "math"

type TmpType bool

type TreeNode struct {
	HasToy TmpType
	Left   *TreeNode
	Right  *TreeNode
}

func (node *TreeNode) FillNode(toys ...TmpType) {
	for _, elem := range toys {
		push(elem, node)
	}
}

func AreToysBalanced(root *TreeNode) bool {
	return (countBool(root.Left) - countBool(root.Right)) == 0
}

func countBool(node *TreeNode) int {
	if node == nil {
		return 0
	}
	count := 0
	count = countBool(node.Left)
	count = countBool(node.Right)
	if node.HasToy == false {
		return count
	}
	return count + 1
}

func push(elem TmpType, node *TreeNode) {
	if node.Left == nil {
		node.Left = &TreeNode{
			Left:   nil,
			Right:  nil,
			HasToy: elem,
		}
	} else if node.Right == nil {
		node.Right = &TreeNode{
			Left:   nil,
			Right:  nil,
			HasToy: elem,
		}
	} else if Height(node.Left) <= Height(node.Right) {
		push(elem, node.Left)
	} else if Height(node.Left) > Height(node.Right) {
		push(elem, node.Right)
	}
}

func Height(node *TreeNode) float64 {
	if node == nil {
		return 0
	}
	return math.Max(Height(node.Left), Height(node.Right)) + 1
}
