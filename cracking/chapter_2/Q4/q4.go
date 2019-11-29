/*
Partion:  Write code to partion a linked list around a value x, such that all node less than x
come before all node greater than equal to x. If x is contained within the list, all values of x only need to
be after the elements less than x. x can appear anywhere in the right partition. It does not need to appear between
the left and right partition.
*/

package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func main() {
	//create a sample LL
	test := []int{3, 5, 8, 5, 10, 2, 1}
	head := &Node{Value: 3, Next: nil}
	temp := head
	for i := 1; i < len(test); i++ {
		newNode := &Node{Value: test[i], Next: nil}
		temp.Next = newNode
		temp = temp.Next
	}
	printList(head)
	printList(partition(5, head))
}

func partition(x int, head *Node) *Node {
	var leftSide, rightSide *Node = nil, nil

	temp := head
	for temp != nil {
		if temp.Value < x {
			leftSide = addNode(leftSide, temp.Value)
		} else {
			rightSide = addNode(rightSide, temp.Value)
		}
		temp = temp.Next
	}

	iter := leftSide
	for iter.Next != nil {
		iter = iter.Next
	}
	iter.Next = rightSide

	return leftSide
}

func addNode(list *Node, value int) *Node {
	if list == nil {
		list = &Node{Value: value, Next: nil}
		return list
	} else {
		temp := &Node{Value: value, Next: list}
		return temp
	}
}

// helper

func printList(h *Node) {
	temp := h
	print("[")
	for temp != nil {
		fmt.Printf("%v ", temp.Value)
		if temp.Next != nil {
			fmt.Print("-> ")
		}
		temp = temp.Next
	}
	print("]\n")
}
