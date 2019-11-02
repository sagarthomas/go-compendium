package main

import "fmt"

func main() {
	//fmt.Println(fib(34))

	var v, i, sum int
	for v < 4000000 {

		if v%2 == 0 {
			sum += v
		}
		fmt.Printf("%v: %v\n", i, v)
		i++
		v = fib(i)
	}
	fmt.Println(i - 1)
	fmt.Println(sum)
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
