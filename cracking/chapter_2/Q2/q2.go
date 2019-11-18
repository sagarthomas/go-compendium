/*
Return  kth to last: Implement an algorithm to find the kth to last elements of a singly linked list
*/

package main

type Node struct {
	Value int
	Next  *Node
}

// Assumption that k = 1 means last element
func main() {
	head := &Node{Value: 0, Next: nil}
	temp := head
	for i := 1; i < 6; i++ {
		newNode := &Node{Value: i, Next: nil}
		temp.Next = newNode
		temp = temp.Next
	}
	result := make(chan int)
	go findKthToLast(3, head, result)
	println("Answer: ")
	println(<-result) // Waits till it recieves the result

}

// Recursive soltuion using a go channel. The function will keep calling itself until it reaches the end
// of the list. At this point, it will pass back a counter that increments on each pass. When the counter == k,
// We send the current node's value through the go channel back to main()
func findKthToLast(k int, head *Node, result chan int) int {
	if head == nil {
		return 0
	}
	counter := findKthToLast(k, head.Next, result)
	counter++
	if counter == k {
		result <- head.Value
	}
	return counter
}
