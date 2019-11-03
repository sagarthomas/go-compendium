/*
Is Unique: Implement and algorithm to determine if a string has all unique characters.
What if you cannot use additional data structures?
*/
package main

import "fmt"

func main() {
	fmt.Println(isUnique("hello"))
	fmt.Println(isUnique("race"))

	fmt.Println(isUnique2("hello"))
	fmt.Println(isUnique2("race"))
}

// Implementation with a map: N
func isUnique(input string) bool {
	m := make(map[rune]bool)
	for _, c := range input {
		if m[c] == true {
			return false
		}
		m[c] = true
	}
	return true
}

// Implementation with no data structures: N^2
func isUnique2(input string) bool {
	for i := range input {
		for k := range input {
			if input[i] == input[k] && i != k {
				return false
			}
		}
	}
	return true
}
