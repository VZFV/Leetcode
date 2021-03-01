package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{val: 0}				// new linked list with head
	num1, num2, carry, current := 0,0,0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil{
			num1 = 0
		} else {
			num1 = l1.val
			l1 = l1.Next
		}
		if l2 == nil {
			num2 = 0
		} else {
			num2 = l2.val
			l2 = l2.Next
		}
		current.Next = &ListNode{val: (num1+num2+carry)%10}
		current = current.Next
		carry = (num1 + num2 + carry) / 10
	}
	return head.Next
}
