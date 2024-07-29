package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil && n == 1 {
		return nil
	}

	var curr *ListNode
	var prev *ListNode
	probe := head
	nodeCount := 0

	for probe != nil {
		probe = probe.Next
		nodeCount++

		if nodeCount == n {
			curr = head
		}
		if nodeCount > n {
			curr = curr.Next
			if nodeCount == n+1 {
				prev = head
			} else {
				prev = prev.Next
			}
		}
	}

	if curr == head {
		return curr.Next
	}

	prev.Next = curr.Next

	return head
}

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	removeNthFromEnd(node, 2)
}
