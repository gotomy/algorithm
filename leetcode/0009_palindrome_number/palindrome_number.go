package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(1000))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 9 {
		return true
	}

	if x%10 == 0 {
		return false
	}

	rightPart := 0
	for x > rightPart {
		rightPart = rightPart*10 + x%10
		x = x / 10
	}

	return x == rightPart || x == rightPart/10
}
