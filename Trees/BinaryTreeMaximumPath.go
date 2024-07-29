package trees

import "math"

var maxPath int

func maxPathSum(root *TreeNode) int {
	maxPath = math.MinInt
	findMax(root)
	return maxPath
}

func findMax(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftMax := findMax(root.Left)
	rightMax := findMax(root.Right)
	maxPath = max(maxPath, leftMax+rightMax+root.Val)
	return max(max(leftMax+root.Val, rightMax+root.Val), 0)
}
