package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	l2 := &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 7}}}

	addTwoNumbers(l1, l2)
}

// // The while loop solution
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	carry := 0
	current := dummy

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum = l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
		carry = sum / 10
	}
	return dummy.Next
}

// Recursive solution
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	return sumTwoListNode(l1, l2, 0)
// }

// func sumTwoListNode(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
// 	if l1 == nil && l2 == nil && carry == 0 {
// 		return nil
// 	}

// 	sum := carry
// 	current := &ListNode{}

// 	if l1 != nil {
// 		sum += l1.Val
// 		l1 = l1.Next
// 	}

// 	if l2 != nil {
// 		sum += l2.Val
// 		l2 = l2.Next
// 	}

// 	current.Val = sum % 10
// 	current.Next = sumTwoListNode(l1, l2, sum/10)
// 	return current
// }
