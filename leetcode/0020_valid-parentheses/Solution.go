package main

import "fmt"

func main() {
	fmt.Println(isValid("[{}{}{}]"))
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0)
	for _, item := range s {
		if item == '(' || item == '[' || item == '{' {
			stack = append(stack, item)
		} else if (item == ']' && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			(item == ')' && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			(item == '}' && len(stack) > 0 && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
