package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println("下标为", twoSum(nums, 9))
}

func twoSum(nums []int, target int) []int {
	findMap := make(map[int]int)
	for idx, val := range nums {
		other := target - val
		if v, ok := findMap[other]; ok {
			return []int{v, idx}
		}
		findMap[val] = idx
	}
	return nil
}
