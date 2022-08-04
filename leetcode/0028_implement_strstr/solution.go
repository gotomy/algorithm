package main

import "fmt"

func main() {
	fmt.Println(strStr("hello", "ado"))
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	left := 0
	right := left + len(needle)
	for {
		substr := haystack[left:right]
		if substr == needle {
			return left
		} else {
			left = left + 1
			right = right + 1
		}

		if right == len(haystack)+1 {
			return -1
		}
	}
}
