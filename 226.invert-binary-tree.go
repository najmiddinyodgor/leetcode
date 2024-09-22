package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	temp := InvertTree(root.Left)
	root.Left = InvertTree(root.Right)
	root.Right = temp

	return root
}
