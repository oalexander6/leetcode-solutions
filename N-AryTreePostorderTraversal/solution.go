package narytreepostordertraversal

// https://leetcode.com/problems/n-ary-tree-postorder-traversal/submissions/1241198776

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	if len(root.Children) == 0 {
		return []int{root.Val}
	}

	result := make([]int, 0)

	for _, node := range root.Children {
		result = append(result, postorder(node)...)
	}

	return append(result, root.Val)
}
