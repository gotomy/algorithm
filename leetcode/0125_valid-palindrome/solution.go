package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("0P"))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	head, end := 0, len(s)-1
	runes := []rune(s)
	for head < end {
		head_str := runes[head]
		end_str := runes[end]
		if !isLetter(head_str) {
			head++
			continue
		}
		if !isLetter(end_str) {
			end--
			continue
		}

		fmt.Println(head, end, string(head_str), string(end_str))
		if isLetter(head_str) && isLetter(end_str) && head_str == end_str {
			head++
			end--
		} else {
			return false
		}
	}
	return true
}

func isLetter(ch rune) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
