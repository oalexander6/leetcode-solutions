package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}

func getMinimumDifference(root *TreeNode) int {
	leftMin := math.MaxInt
	rightMin := math.MaxInt

	if root.Left != nil {
		leftMin = min(getMinimumDifference(root.Left), abs(root.Val-root.Left.Val))
	}
	if root.Right != nil {
		rightMin = min(getMinimumDifference(root.Right), abs(root.Val-root.Right.Val))
	}

	return min(leftMin, rightMin)
}

func main() {
	input := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   3,
			Right: nil,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Left: nil,
	}

	_ = getMinimumDifference(input)
}
