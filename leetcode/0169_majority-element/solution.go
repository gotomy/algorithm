package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

func majorityElement(nums []int) int {
	numsMap := make(map[int]int)
	for _, v := range nums {
		count, ok := numsMap[v]
		if ok {
			numsMap[v] = count + 1
		} else {
			numsMap[v] = 1
		}
	}

	majorityCount := len(nums) / 2

	for k, v := range numsMap {
		if v > majorityCount {
			return k
		}
	}

	return 0
}
