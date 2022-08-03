package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(2))
}

func generateParenthesis(n int) []string {
	arr := make([]string, 0)
	dfs(&arr, "", n, n)
	return arr
}

func dfs(arr *[]string, s string, left int, right int) {
	if left == 0 && right == 0 {
		*arr = append(*arr, s)
		return
	}

	if left > 0 {
		dfs(arr, s+"(", left-1, right)
	}

	if right > left {
		dfs(arr, s+")", left, right-1)
	}
}
