package tree

import "math"

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func (node *TreeNode) FillNode(toys ...bool) {
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
	count := 1
	count += countBool(node.Left)
	count += countBool(node.Right)
	if node.HasToy == false {
		return 0
	}
	return count
}

func push(elem bool, node *TreeNode) {
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
	} else if height(node.Left) <= height(node.Right) {
		push(elem, node.Left)
	} else if height(node.Left) > height(node.Right) {
		push(elem, node.Right)
	}
}

func height(node *TreeNode) float64 {
	if node == nil {
		return 0
	}
	return math.Max(height(node.Left), height(node.Right)) + 1
}
