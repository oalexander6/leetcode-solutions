package main

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

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	prev := head
	curr := head.Next
	prev.Next = nil

	for curr != nil {
		newCurr := curr.Next
		curr.Next = prev
		prev = curr
		curr = newCurr
	}

	return prev
}
