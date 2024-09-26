package main

import "fmt"

func main() {
	/**
	 	[
			"MyLinkedList",[],
			"addAtHead",[1],
			"addAtTail",[3],
			"addAtIndex",[1,2],
			"get",[1],
			"deleteAtIndex",[0],
			"get",[0],

			"addAtHead",[1],
			"addAtIndex",[3,0],
			"addAtHead",[6],
			"addAtTail",[4],
			"addAtHead",[4],
			"addAtIndex",[5,0],
			"addAtHead"[6],
		]
	*/

	obj := Constructor()
	obj.AddAtTail(1)
	fmt.Println(obj.Get(0))
}
