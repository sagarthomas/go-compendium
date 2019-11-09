/*
Project Euler - Problem #4: Largest palindrome product
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// The number range to look for is 100 - 999 however, it is easier to iterate starting from 999
	// since we are looking for the largest palindrome product. Also, we can stop iterating at 900
	// to prevent looking at lower valued palindromes
	start := time.Now()
	for i := 999; i > 900; i-- {
		for j := 999; j > 900; j-- {
			if isPalindromeNumber(i * j) {
				fmt.Printf("%v * %v : %v\n", i, j, i*j)
				println(time.Since(start).Seconds())
				return
			}
		}
	}
}

func isPalindromeNumber(value int) bool {
	digits := make([]int, 0)
	temp := value

	for temp != 0 {
		digits = append(digits, temp%10)
		temp = temp / 10
	}

	for i := 0; i < len(digits); i++ {
		if digits[i] != digits[len(digits)-1-i] {
			return false
		}
	}
	return true
}
