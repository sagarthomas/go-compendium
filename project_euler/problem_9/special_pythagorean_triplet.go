/*
Project Euler - Special Pythagorean triplet:
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	var a, b, c = 0, 1, 2
	for a = 0; a < b; a++ {
		for b = a + 1; b < c; b++ {
			for c = b + 1; c < 1000; c++ {
				if a*a+b*b == c*c && a+b+c == 1000 {
					fmt.Printf("%d, %d, %d\n", a, b, c)
					fmt.Println(a * b * c)
					fmt.Printf("in %vs\n", time.Since(start).Seconds())
					return
				}
			}
		}
	}

}
