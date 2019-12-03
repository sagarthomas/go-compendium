package main

// file used to test

import "fmt"

func main() {
	//queueTest()
	bagTest()
}

func stackTest() {
	s := Stack{}

	s.Push(1)
	s.Push(2)
	fmt.Println(s.Size())
	fmt.Println(s.Pop()) //should be 2
	s.Pop()
	fmt.Println(s.IsEmpty())
}

func queueTest() {
	q := Queue{}

	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q.IsEmpty()) // false
	fmt.Println(q.Size())    // should be 2
	fmt.Println(q.Dequeue()) // should be 1
	q.Dequeue()
	fmt.Println(q.IsEmpty()) //true
}

func bagTest() {
	b := Bag{}

	fmt.Printf("Empty: %v\n", b.IsEmpty())
	fmt.Printf("Size: %d\n", b.Size())
	b.Add(1)
	b.Add(2)
	b.Add(3)
	for r := range b.Iter() {
		fmt.Println(r)
	}
	fmt.Printf("Size: %d\n", b.Size())
	fmt.Printf("Empty: %v\n", b.IsEmpty())

}
