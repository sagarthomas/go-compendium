/*
Project Euler - Problem #5: Smallest multiple
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import "fmt"

// Things we know:
// 1. All numbers are divisible by 1 so we don't need to waste time computing that
// 2. even numbers are divisible by 2 etc
// 3. if some value b is divisble by a, b is also divisble by the factors of a. ex. 40 % 20 == 0 -> 40 % (10, 5, 4, 2, 1) == 0
func main() {

	// For each number from 1 -> 20 we can create a list of all factors of that number for faster lookup time later on
	// factorTable := make(map[int][]int)
	// for i := 1; i <= 20; i++ {
	// 	factorTable[i] = findFactors(i)
	// }
	current := 20
	failedDivisor := make(chan int)
	for {
		go checkMultiples(current, failedDivisor)
		x := <-failedDivisor
		if x == 0 {
			break
		}
		current += 20 // increment by 20 since we don't care about values that aren't divisible by 20

		fmt.Println("======")
	}

	fmt.Println(current)
}

func findFactors(i int) []int {
	// There should be less factors than the sqrt of i but we'll implement that later
	result := make([]int, 0)
	for j := i; j > 0; j-- {
		if i%j == 0 {
			result = append(result, j)
		}
	}
	fmt.Printf("factors of %v: %v\n", i, result)
	return result
}

func checkMultiples(v int, c chan int) bool {
	// create a simple int-bool map for traking values 1-20

	for i := 2; i < 20; i++ {
		if v%i != 0 {
			fmt.Printf("%v is not divisble by %v\n", v, i)
			c <- i
			return false
		}
	}
	c <- 0
	return true
}
