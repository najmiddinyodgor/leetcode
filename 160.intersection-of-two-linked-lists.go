package main

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := getLength(headA), getLength(headB)

	for lenA > lenB {
		lenA--
		headA = headA.Next
	}

	for lenB > lenA {
		lenB--
		headB = headB.Next
	}

	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}

		headA = headA.Next
		headB = headB.Next
	}

	return nil
}

func getLength(node *ListNode) int {
	l := 0

	for node != nil {
		node = node.Next
		l++
	}

	return l
}
