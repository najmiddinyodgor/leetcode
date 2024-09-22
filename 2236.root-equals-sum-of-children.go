package main

func CheckTree(root *TreeNode) bool {
	return root.Val == (root.Left.Val + root.Right.Val)
}
