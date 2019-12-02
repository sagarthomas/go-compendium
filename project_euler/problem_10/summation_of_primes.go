/*
Project Euler - Summation of primes
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/
package main

import (
	"fmt"
	"math"
	"time"
)

/*
We can use the Sieve of Eratosthenes to generate a list of primes up to 2 million
*/
func main() {
	start := time.Now()
	primes := sieve(2, 2000000)
	sum := 0
	for i, b := range primes {
		if b == false && i > 1 { // 1 is not prime
			sum += i
		}
	}

	fmt.Printf("The sum of all primes up to 2 million is: %d in %vs\n", sum, time.Since(start).Seconds())
}

func sieve(start int, n int) []bool {
	list := make([]bool, n)
	for i := start; i < int(math.Sqrt(float64(n))); i++ {
		if list[i] == true {
			continue
		}
		for r := i * i; r < n; r += i {
			list[r] = true
		}
	}
	return list
}
