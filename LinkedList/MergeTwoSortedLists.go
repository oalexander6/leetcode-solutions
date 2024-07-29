package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var primaryHead *ListNode
	var primaryPtr *ListNode
	var secondaryPtr *ListNode

	if list1.Val > list2.Val {
		primaryHead = list2
		primaryPtr = list2
		secondaryPtr = list1
	} else {
		primaryHead = list1
		primaryPtr = list1
		secondaryPtr = list2
	}

	for {
		if secondaryPtr == nil {
			break
		}

		if primaryPtr.Next == nil || secondaryPtr.Val < primaryPtr.Next.Val {
			nextSecondary := secondaryPtr.Next

			secondaryPtr.Next = primaryPtr.Next
			primaryPtr.Next = secondaryPtr

			primaryPtr = primaryPtr.Next
			secondaryPtr = nextSecondary
		} else {
			if primaryPtr.Next == nil {
				primaryPtr.Next = secondaryPtr
				break
			}

			primaryPtr = primaryPtr.Next
		}
	}

	return primaryHead
}
