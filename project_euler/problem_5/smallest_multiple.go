/*
Project Euler - Problem #5: Smallest multiple
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import (
	"fmt"
	"time"
)

// Originally implemented brute force by iterating by 20 and checking if the value is evenly divisible by 2..19
// This can be further optimized by checking numbers 11..19 only. This is because -> if a % b == 0, then a % factors(b) == 0
// Since the range 11..19 contains numbers that are composite values from 1..10, a value being divisible by 11..19 is divisible
// by 1..10.
// To optimize even further, We notice that we are essentially looking for the LCM of numbers 1..20 and we can do so by
// using the forumula LCM(a, b) = a * b/ gcd(a, b) which can be extended to LCM(a, b, c) = LCM(a, LCM(b,c)), where gcd is greatest
// common divisor
func main() {
	start := time.Now()
	a := make([]int, 20)
	for i := 0; i < 20; i++ {
		a[i] = i + 1
	}

	// fmt.Println(a)
	fmt.Printf("Smallest positive number is %v in %vs", multiLCM(a), time.Since(start).Seconds())
}

func multiLCM(values []int) int {
	if len(values) == 1 {
		return values[0]
	}
	return lcm(values[0], multiLCM(values[1:]))
}

func lcm(a int, b int) int {
	return (a * b) / gcd(a, b)
}

func gcd(a int, b int) int {
	for b != 0 {
		temp := a
		a = b
		b = temp % b
	}
	return a
}
