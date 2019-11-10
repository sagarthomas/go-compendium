/*
String Compression: Implement a method to perform basic string compression using the counts of
repeated characters. Example: aabcccccaaa -> a2b1c5a3. If the compressed string is would not become
shorter than the orginal string, return the original string. String will only contain lowecase and
uppercase (a - z)
*/

package main

import (
	"fmt"
	"strconv"
)

type char struct {
	value rune
	count int
}

// Can implement a hash table but need to take into account the location of the characters
// ex. aabccccaaa should be a2b1c4a3 instead of a5b1c4
func main() {

	fmt.Println(compress("aabccccaaa"))  // return a2b1c4a3
	fmt.Println(compress("aaggbcccgca")) // return original
	fmt.Println(compress("aba"))         // should return orignal string
}

func compress(input string) string {
	charMap := make([]char, 0)
	prev := rune(' ') // since input will always be a - z, we can safely use this as the first prev entry
	i := 0
	for _, v := range input {
		if v != prev {
			i++
			charMap = append(charMap, char{value: v, count: 1})
			prev = v
		} else {
			charMap[i-1].count++
		}
	}
	// The compressed string will always be the number of entries in the charMap* 2 (each character has a number of occurances )
	if len(charMap)*2-1 >= len(input) {
		return input
	}
	result := ""
	for _, c := range charMap {
		result += string(c.value) + strconv.Itoa(c.count)
	}

	return result
}
