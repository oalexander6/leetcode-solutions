package leafsimilartrees

import "slices"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1 := dfsLeaves(root1)
	leaves2 := dfsLeaves(root2)

	return (slices.Compare(leaves1, leaves2) == 0)
}

func dfsLeaves(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	leaves := append(dfsLeaves(root.Left), dfsLeaves(root.Right)...)

	if len(leaves) > 0 {
		return leaves
	} else {
		return []int{root.Val}
	}
}
