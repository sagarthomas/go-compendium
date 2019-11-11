package main

/*
FIFO Queue
*/
type Queue struct {
	head *Node
}

func (q *Queue) Enqueue(item interface{}) {
	if q.head == nil {
		q.head = &Node{Value: item, Next: nil}
	} else {
		q.head.Next = &Node{Value: item, Next: nil}
	}
}

func (q *Queue) Dequeue() interface{} {
	temp := q.head
	q.head = temp.Next
	return temp.Value
}

func (q Queue) IsEmpty() bool {
	if q.Size() == 0 {
		return true
	}
	return false
}

func (q Queue) Size() int {
	if q.head == nil {
		return 0
	}
	iter := q.head
	count := 0
	for iter != nil {
		count++
		iter = iter.Next
	}
	return count
}
