/*
Palindrome Permutation: Given a string, write a function to check if it is a permutation of a palindrome.
A palindrome is a word that is the same forwards or backwards. A permutation is a rearrangment of letters.
The palindrome does not need to be limited to just dictionary words.
*/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Printf("%v\n", isPalindrome("Tact Coa"))
	fmt.Printf("%v\n", isPalindrome("log log"))
	fmt.Printf("%v\n", isPalindrome("Tact Coae"))
	fmt.Printf("%v\n", isPalindrome("cocoppo"))

}

func isPalindrome(input string) bool {
	// One property of a palindrome is that we need pairs for each character except for a middle character
	// in the case that the number of none-space characters is odd
	// An assumption made looking at the example "Tact coa" is that upper case does not matter
	// The positioning of spaces in a string also does not seem to impact how the palindrome is read
	input = strings.ToLower(input)
	table := make(map[rune]int) // create a hash table to count character occurances
	specialCount := 0
	for _, r := range input {
		if !unicode.IsLetter(r) {
			specialCount++
		} else {
			table[r]++
		}
	}
	isEvenLength := ((len(input)-specialCount)%2 == 0)
	exception := true // odd count character exception

	for _, v := range table {
		if v%2 != 0 && isEvenLength {
			return false
		}

		if v%2 != 0 && !isEvenLength && !exception {
			return false
		}

		if v%2 != 0 && !isEvenLength && exception {
			exception = false
		}
	}
	return true
}
