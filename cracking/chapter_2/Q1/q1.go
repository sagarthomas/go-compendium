/*
Remove dups: Write code to remove duplicates from an unsorted Linkedlist.
*/

package main

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

// Solution using a hash table. If no extra memory is allowed to be used
// we can either compare each node to the rest of the list and find duplicates O(n^2)
// or sort the list and do a side by side comparison O(n) or use two pointers and for each
// value, use the second pointer to look for duplicates and remove them
func main() {
	head := &Node{1, nil}
	head.Next = &Node{1, nil}
	head.Next.Next = &Node{3, nil}
	head.Next.Next.Next = &Node{1, nil}

	printList(head)
	head = removeDups(head)
	printList(head)
}

func removeDups(head *Node) *Node {
	table := make(map[interface{}]bool)
	temp := head
	prev := head
	for temp != nil {
		if table[temp.Value] {
			// Found dup
			prev.Next = temp.Next
		}
		table[temp.Value] = true
		prev = temp
		temp = temp.Next
	}
	return head
}

func printList(head *Node) {
	temp := head
	for temp != nil {
		fmt.Println(temp.Value)
		temp = temp.Next
	}
	fmt.Println("===")
}
