package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("239"))
}

var combinations []string
var letterMapping = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := letterMapping[digit]
		lettersLen := len(letters)
		for i := 0; i < lettersLen; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}
