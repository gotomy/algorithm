package main

import "fmt"

func main() {
	s := "a"
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	right := len(s)
	left := right
	for {
		fmt.Println("left:", left, ",right:", right, ",", s[left:right])

		rightStr := s[right-1 : right]
		if rightStr == " " {
			right--
			left = right
		} else {
			leftStr := s[left-1 : left]
			if leftStr != " " {
				left--
			} else {
				break
			}
		}
		if left == 0 {
			break
		}
	}
	return right - left
}
