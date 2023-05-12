package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}

func containsDuplicate(nums []int) bool {
	counts := make(map[int]int)
	for _, val := range nums {
		if _, ok := counts[val]; ok {
			return true
		} else {
			counts[val] = 1
		}
	}

	return false
}
