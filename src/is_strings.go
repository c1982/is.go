package Is

import (
	"strings"		
)

//Lowercase check text is lower of the string
func Lowercase(input string) bool {
	return strings.ToLower(input) == input
}

//Uppercase check capital letter of the string
func Uppercase(input string) bool {
	return strings.ToUpper(input) == input
}

//EndWith check end string of the string
func EndWith(input, suffix string) bool {
	return len(input) >= len(suffix) && input[len(input)-len(suffix):] == suffix
}

//StartWith check start string of the string
func StartWith(input, prefix string) bool {
	return len(input) >= len(prefix) && input[0:len(prefix)] == prefix
}

//Capitalized check start capital letter of the string
func Capitalized(input string) bool {
	var words = strings.Split(input, " ")

	for i := 0; i < len(words); i++ {
		word := words[i]
		if len(word) > 0 {
			chr := word[0:1]			
			if !Uppercase(chr) {
				return false
			}
		}
	}

	return true
}

//Palindrome is a given string palindrome?
func Palindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}
