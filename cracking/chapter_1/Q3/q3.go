/*
URLify: Write a method to replaces all spaces in a string with '%20'. You may assume that the string
has sufficient space at the end to hold additional characters, and you are given the length of the 'true'
string.
*/
package main

import "fmt"

func main() {
	fmt.Print("%v", URLify([]rune("Mr John Smith    "), 13)) // Used rune array to perform in place
	fmt.Print("%v", URLify([]rune("Mr John     Smith"), 13)) // Used rune array to perform in place

}

func URLify(input []rune, length int) []rune {
	//count the number of spaces in the 'true' string
	spaceCount := 0
	for i := 0; i < length; i++ {
		if input[i] == ' ' {
			spaceCount++
		}
	}

}

func shift(input []rune, shiftAmount int) {
	if input[0] == ' ' && shiftAmount == 0 {
		return
	}
	shift(input[1:])
	// input[1] is now a space due to shift operation
	input[1] = input[0]
	input[0] = ' '
	return
}
