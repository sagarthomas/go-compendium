package main

// file used to test

import "fmt"

func main() {
	//queueTest()
	//bagTest()
	UFTest()
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

func UFTest() {
	uf := NewUF(10)
	fmt.Printf("Number of components: %d\n", uf.Count())
	uf.Union(3, 4)
	fmt.Printf("%d is connected to root : %d\n", 4, uf.Find(4))
	fmt.Println(uf.Connected(3, 4)) // true
	uf.Union(4, 5)
	fmt.Printf("%d is connected to root : %d\n", 3, uf.Find(3))
	uf.Union(3, 9)
	fmt.Println(uf.Connected(9, 5)) // true
	for i := 0; i < 10; i++ {
		fmt.Printf("%d is connected to root : %d\n", i, uf.Find(i))
	}

	fmt.Printf("Number of components: %d\n", uf.Count()) // 7
}
