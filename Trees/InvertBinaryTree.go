package trees

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	invertTree(root.Right)
	invertTree(root.Left)

	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	return root
}
