package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	fmt.Println(fourSum(nums, 0))
}

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	var ret [][]int
	for i := 0; i < len(nums)-3; i++ {
		n1 := nums[i]

		if i > 0 && n1 == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			n2 := nums[j]
			if j > i+1 && n2 == nums[j-1] {
				continue
			}
			l := j + 1
			r := len(nums) - 1
			for l < r {
				n3 := nums[l]
				n4 := nums[r]
				sum := n1 + n2 + n3 + n4
				if sum < target {
					l++
				} else if sum > target {
					r--
				} else {
					ret = append(ret, []int{n1, n2, n3, n4})
					for l < r && n3 == nums[l+1] {
						l++
					}
					for l < r && n4 == nums[r-1] {
						r--
					}
					r--
					l++
				}
			}
		}
	}
	return ret
}
