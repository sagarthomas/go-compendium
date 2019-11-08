/*
One Away: There are three types of edits that can be performed on a string: Inserta character,
Remove a character, replace a character. Given two strings, write a function to check if they are
one edit away (or zero edits).

*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(isOneAway("pale", "ple"))    // true - deletion
	fmt.Println(isOneAway("pales", "pale"))  // true - insertion
	fmt.Println(isOneAway("pale", "bale"))   // true - replacement
	fmt.Println(isOneAway("pale", "pale"))   // true - zero edit
	fmt.Println(isOneAway("pale", "bake"))   // false - 2 replacement
	fmt.Println(isOneAway("pale", "paleee")) // false - 2 insertion
	fmt.Println(isOneAway("late", "te"))     // false - 2 deletion
	fmt.Println(isOneAway("abcd", "adbee"))  // false - subtituion + insertion

}

func isOneAway(s1 string, s2 string) bool {
	if (len(s1) - len(s2)) == 1 {
		return diffCheck(s1, s2)
	} else if (len(s2) - len(s1)) == 1 {
		return diffCheck(s2, s1)
	} else if len(s1) == len(s2) {
		return diffCheck(s1, s2)
	} else if s1 == s2 {
		return true
	} else {
		return false
	}
}

//helper
func diffCheck(s1 string, s2 string) bool {
	diffChar := 0
	offset := 1
	if len(s1) == len(s2) {
		offset = 0
	}
	for i := 0; i < len(s1)-offset; i++ {
		if diffChar > 1 {
			return false
		}
		if s1[i] != s2[i] {
			diffChar++
		}
	}
	return true
}
