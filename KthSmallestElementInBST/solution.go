package kthsmallestelementinbst

// https://leetcode.com/problems/kth-smallest-element-in-a-bst/submissions/1252214223

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	values := make([]int, 0)

	values = inorder(values, root)

	return values[k-1]
}

func inorder(results []int, node *TreeNode) []int {
	if node == nil {
		return results
	}

	results = inorder(results, node.Left)
	results = append(results, node.Val)
	results = inorder(results, node.Right)

	return results
}
