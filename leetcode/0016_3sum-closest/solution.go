package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		start := i + 1
		end := len(nums) - 1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if math.Abs(float64(target)-float64(ans)) > math.Abs(float64(target)-float64(sum)) {
				ans = sum
			}
			if sum > target {
				end = end - 1
			} else if sum < target {
				start = start + 1
			} else {
				return sum
			}
		}
	}
	return ans
}
