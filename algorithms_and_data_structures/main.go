package main

// file used to test

import "fmt"

func main() {
	s := Stack{}

	s.Push(1)
	s.Push(2)
	fmt.Println(s.Size())
	fmt.Println(s.Pop()) //should be 2
	s.Pop()
	fmt.Println(s.IsEmpty())
}
