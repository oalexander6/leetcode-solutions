package hascycle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slowPtr := head
	fastPtr := head.Next

	for {
		if slowPtr == nil || fastPtr == nil {
			return false
		}
		if slowPtr == fastPtr {
			return true
		}
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next
		if fastPtr != nil {
			fastPtr = fastPtr.Next
		}
	}

	return false
}
