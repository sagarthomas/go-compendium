/*
Project Euler - Largest Prime Factor:
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

package main

import (
	"time"
)

// Any interger greater than 1 is either a prime number, or a composite of prime numbers
// We can find the largest prime using division. Starting with the value i=2, we do the operation
// n % i to determine if n is divisible by i. If it is, we divide n by i and store the result as the new value of n
// and attempt to divide by the same number again until it no longer divides without remainder. If there is a remainder,
// we iterate i by one and repeat the same process until i == n, which would be our greatest prime factor
func main() {
	start := time.Now()
	n := 600851475143
	i := 2
	for i != n {
		if n%i == 0 {
			//fmt.Printf("i: %v\n", i)
			n = n / i
			//fmt.Printf("n: %v\n", n)
		} else {
			i++
		}
	}
	println(n)
	println(time.Since(start).Seconds())
}
