package chapter_2

// Implementing a singly LinkedList

type Node struct {
	Value interface{}
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) Append(value interface{}) {
	if l.Head == nil {
		l.Head.Value = value
		l.Tail = l.Head
		return
	}
	node := &Node{Value: value, Next: nil}
	l.Tail.Next = node
}

// Removes the first occurance of a value in the LinkedList
func (l *LinkedList) Remove(value interface{}) {

}
