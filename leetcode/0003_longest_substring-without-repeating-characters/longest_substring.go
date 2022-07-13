package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int)
	max := 0
	left := 0
	for i := 0; i < len(s); i++ {
		if val, ok := mp[s[i]]; ok {
			left = maxVal(left, val+1)
		}
		mp[s[i]] = i
		max = maxVal(max, i-left+1)
	}
	return max
}

func maxVal(x, y int) int {
	if x < y {
		return y
	}

	return x
}
