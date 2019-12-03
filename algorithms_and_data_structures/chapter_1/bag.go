package main

type Bag struct {
	head *Node
	size int
}

func (b *Bag) Add(item interface{}) {
	if b.head == nil {
		b.head = &Node{Value: item, Next: nil}
		b.size++
		return
	}
	temp := b.head
	b.head = &Node{Value: item, Next: temp}
	b.size++
}

func (b Bag) IsEmpty() bool {
	if b.size == 0 {
		return true
	}
	return false
}

func (b Bag) Size() int {
	return b.size
}

func (b Bag) Iter() <-chan interface{} {
	channel := make(chan interface{})
	go func() {
		temp := b.head
		for temp != nil {
			channel <- temp.Value
			temp = temp.Next
		}
		close(channel) // so the client won't be waiting forever
	}()
	return channel
}
