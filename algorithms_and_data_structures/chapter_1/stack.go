package main

type Stack struct {
	head *Node
}

func (s *Stack) Push(item interface{}) {
	if s.head == nil {
		s.head = &Node{Value: item, Next: nil}
		return
	}
	temp := s.head
	s.head = &Node{Value: item, Next: temp}
}

func (s *Stack) Pop() interface{} {
	temp := s.head
	s.head = s.head.Next
	return temp.Value
}

func (s Stack) IsEmpty() bool {
	if s.Size() == 0 {
		return true
	}
	return false
}

func (s Stack) Size() int {
	if s.head == nil {
		return 0
	}
	temp := s.head
	count := 0
	for temp != nil {
		count++
		temp = temp.Next
	}
	return count
}
