package test

import (
	run "day05/ex00/tree"
	"testing"
)

func TestTrue1(t *testing.T) {
	node := run.TreeNode{}
	node.FillNode(true, true, false, false)
	isSame := run.AreToysBalanced(&node)
	if isSame != true {
		t.Errorf("branches are not equal, return value: %t", isSame)
	}
}

func TestTrue2(t *testing.T) {
	node := run.TreeNode{}
	node.FillNode(true, true, true, true)
	isSame := run.AreToysBalanced(&node)
	if isSame != true {
		t.Errorf("branches are not equal, return value: %t", isSame)
	}
}

func TestFalse1(t *testing.T) {
	node := run.TreeNode{}
	node.FillNode(true, false, true, false)
	isSame := run.AreToysBalanced(&node)
	if isSame == true {
		t.Errorf("branches are not equal, return value: %t", isSame)
	}
}

func TestFalse2(t *testing.T) {
	node := run.TreeNode{}
	node.FillNode(true, false, true, true, false)
	isSame := run.AreToysBalanced(&node)
	if isSame == true {
		t.Errorf("branches are not equal, return value: %t", isSame)
	}
}
