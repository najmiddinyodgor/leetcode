package main

func InsertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	n1 := head
	n2 := n1.Next

	for n1.Next != nil {
		n := ListNode{Val: getCommonAdvisor(n1.Val, n2.Val), Next: n2}

		n1.Next = &n

		n1 = n2
		n2 = n1.Next
	}

	return head
}

func getCommonAdvisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
