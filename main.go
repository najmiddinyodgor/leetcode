package main

import "fmt"

func main() {
	node1 := &ListNode{1, nil}
	node2 := &ListNode{2, nil}
	node3 := &ListNode{3, nil}
	node4 := &ListNode{4, nil}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node1

	fmt.Println(HasCycle(node1))
}
