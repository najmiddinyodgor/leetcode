package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 3, 5, 2, 6}
	fmt.Println(RemoveElement(a, 3))
	fmt.Println(a)
}
