package main

import "fmt"

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value int
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{
		value: value,
	}
}

func (n *TreeNode) insertLeft(value int) *TreeNode {
	n.left = NewTreeNode(value)
	return n.left
}

func (n *TreeNode) insertRight(value int) *TreeNode {
	n.right = NewTreeNode(value)
	return n.right
}

// depth-first ouptut
func (n *TreeNode) output() {
	if n == nil {
		return
	}
	if n.left != nil {
		n.left.output()
	}

	fmt.Print(n.value, ",")

	if n.right != nil {
		n.right.output()
	}
}

func main() {
	var root *TreeNode

	root = NewTreeNode(50)
	node30 := root.insertLeft(30)
	node80 := root.insertRight(80)
	node20 := node30.insertLeft(20)
	node30.insertRight(40) // node40
	node20.insertLeft(10)  // node10
	node70 := node80.insertLeft(70)
	node90 := node80.insertRight(90)
	node70.insertLeft(60)   // node60
	node90.insertLeft(85)   // node85
	node90.insertRight(100) // node100

	root.output()
	fmt.Println()
	fmt.Println("BST?", checkBST(root))
}

func checkBST(n *TreeNode) bool {
	if n == nil {
		return true
	}

	if n.left != nil && n.left.value > n.value {
		return false
	}

	if n.right != nil && n.right.value < n.value {
		return false
	}

	return checkBST(n.left) && checkBST(n.right)
}
