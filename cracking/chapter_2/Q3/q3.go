/*
Delete middle node: Implement an algorithm to delete a node in the middle of a list (anything between the first and last)
of a singly linked list, given only access to that node
*/

package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func main() {
	head := &Node{Value: 0, Next: nil}
	temp := head
	sample := &Node{}
	for i := 1; i < 6; i++ {
		newNode := &Node{Value: i, Next: nil}
		if i == 3 {
			sample = newNode // we will use sample to test out deleteNode
		}
		temp.Next = newNode
		temp = temp.Next
	}
	printList(head)
	deleteNode(sample)
	printList(head)
}

// We can implement this by deleting the node after n, after copying over the content from n.next
// Copy the value and next pointer of n.next to n and return. Go's garbage collection will remove the old
// n.next as it is no longer in use
func deleteNode(n *Node) {
	if n == nil || n.Next == nil { // edge cases - if the n given is nil or if this is the last element in the list
		return
	}
	temp := n.Next
	n.Value = temp.Value
	n.Next = temp.Next
}

func printList(h *Node) {
	temp := h
	print("[")
	for temp != nil {
		fmt.Printf("%v, ", temp.Value)
		temp = temp.Next
	}
	print("]")
}
