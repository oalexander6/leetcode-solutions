package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type Pair struct {
    Big int
    Small int
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    curr := head

    cache := make(map[Pair]int)

    for curr != nil && curr.Next != nil {
        var valToInsert int
        currPair := Pair{Big: max(curr.Val, curr.Next.Val), Small: min(curr.Val, curr.Next.Val)}

        if val, found := cache[currPair]; found {
            valToInsert = val
        } else {
            valToInsert = gcd(curr.Val, curr.Next.Val)
            cache[currPair] = valToInsert
        }

        newNode := &ListNode{
            Val: valToInsert,
            Next: curr.Next,
        }
        curr.Next = newNode

        curr = newNode.Next
    }

    return head
}

func gcd(a int, b int) int {
    for b != 0 {
        t := b
        b = a % b
        a = t
    }
    return a
}
