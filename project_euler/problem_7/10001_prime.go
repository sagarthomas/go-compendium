/*
Project Euler - 10, 001 prime:
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/

package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

/*
- In order to find all primes within a given range, the sieve of eratosthenes is a good method
- However, we do not know what range the 10,001 prime is in
- One method is to increase the range of N while counting the primes -> failed because increasing the range means we have to start over from 2
for the new range
- sieve optimization #1: we can stop iterating the outer loop at the sqrt(N) as we won't
 find any more primes beyond that. The inner loop can start at the square of the prime found
 - Found a handy theorem called Rosser's theorem which states that the nth prime is found within the range less than n * (ln n + ln ln n)
 more details here: https://en.wikipedia.org/wiki/Prime_number_theorem#Approximations_for_the_nth_prime_number
 - For this case, the runtime was extremely fast but this approch might not work for larger numbers as making a sieve requires to create a boolean
 array of size (n * (ln n + ln ln n))
*/
func main() {
	defer timeTrack(time.Now())
	sieveRange := rosserTheorem(10001)
	count := 0
	for p, b := range sieve(2, sieveRange) {
		if b == false && p >= 2 {
			count++
		}
		if count == 10001 {
			fmt.Printf("Answer %d\n", p)
			return
		}
	}
}

func rosserTheorem(n int) int {
	N := float64(n)
	p := N * (math.Log(N) + math.Log(math.Log(N)))
	return int(math.Ceil(p))
}

func sieve(start, n int) []bool {
	list := make([]bool, n)

	for i := start; i < int(math.Sqrt(float64(n))); i++ {
		if list[i] == true {
			continue
		}
		for j := i * i; j < n; j += i {
			list[j] = true
		}
	}
	return list
}

// timer function
func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}
