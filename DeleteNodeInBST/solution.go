package deletenodeinbst

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		// zero or one children case
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// two children case
		root.Val = minVal(root.Right)
		root.Right = deleteNode(root.Right, root.Val)
	}

	return root
}

func minVal(root *TreeNode) int {
	if root.Left != nil {
		return minVal(root.Left)
	}

	return root.Val
}
