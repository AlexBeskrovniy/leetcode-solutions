package easy

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

var list ListNode

func addValues(l1 *ListNode, l2 *ListNode, reminder int) *ListNode {
	if l1 == nil && l2 == nil {
		return &list
	}

	var rem = 0

	list.Val = l1.Val + l2.Val + reminder
	if list.Val > 9 {
		list.Val = list.Val - 10
		rem = 1
	}

	list.Next = addValues(l1.Next, l2.Next, rem)

	return &list
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addValues(l1, l2, 0)
}
