/*
Check Permutation:  Given two strings, write a method to decide if one string
is a permutation of another.
*/
package main

import "fmt"

func main() {
	fmt.Println(isPermutation("ABCD", "DBAC"))  // yes
	fmt.Println(isPermutation("top", "pop"))    // no
	fmt.Println(isPermutation("chart", "dart")) // no
	fmt.Println(isPermutation("tool", "loot"))  // yes

}

func isPermutation(s1 string, s2 string) bool {
	// In order to be a permutation, must be the same length
	if len(s1) != len(s2) {
		return false
	}

	// We count the number of characters in each string and compare
	s1Count := make(map[rune]int)
	s2Count := make(map[rune]int)

	s1r := []rune(s1)
	s2r := []rune(s2)
	for i := range s1r {
		s1Count[s1r[i]]++
		s2Count[s2r[i]]++
	}

	for v := range s1Count {
		if s2Count[v] != s1Count[v] {
			return false
		}
	}
	return true
}
