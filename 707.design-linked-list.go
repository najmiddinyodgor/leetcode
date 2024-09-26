package main

import (
	"fmt"
	"strconv"
)

type MyLinkedList struct {
	size int
	head *Node
}

type Node struct {
	val  int
	next *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}

	current := this.head

	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.size++
	this.head = &Node{val, this.head}
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{val, nil}

	if this.head == nil {
		this.head = node
	} else {
		current := this.head

		for current.next != nil {
			current = current.next
		}

		current.next = node
	}

	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}

	if index <= 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.size {
		this.AddAtTail(val)
		return
	}

	current := this.head

	for i := 1; i < index; i++ {
		current = current.next
	}

	node := Node{val, current.next}
	node.next = current.next
	current.next = &node
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}

	if index == 0 {
		this.head = this.head.next
	} else {
		current := this.head
		for i := 1; i < index; i++ {
			current = current.next
		}
		current.next = current.next.next
	}

	this.size--
}

func (this *MyLinkedList) Print() {
	result := ""
	next := this.head

	for next != nil {
		result += strconv.Itoa(next.val) + ","

		next = next.next
	}

	fmt.Println(result)
}
