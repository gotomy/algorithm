package main

import "fmt"

func main() {
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak("leetcode", wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	l := len(s)
	wordMap := map[string]bool{}
	for _, v := range wordDict {
		wordMap[v] = true
	}
	dp := make([]bool, l+1)
	dp[0] = true
	for i := 1; i <= l; i++ {
		for j := i - 1; j >= 0; j-- {
			suffix := s[j:i]
			fmt.Println(i, j, suffix, dp[j])
			if wordMap[suffix] && dp[j] {
				dp[i] = true
				break
			}
		}
	}

	return dp[l]
}
