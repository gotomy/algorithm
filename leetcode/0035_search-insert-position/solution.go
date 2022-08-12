package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 7))
}

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	m := (l + r) / 2
	for {
		temp := nums[m]

		if temp < target {
			l = m
			m = (l + r) / 2
		} else if temp > target {
			r = m
			m = (l + r) / 2
		} else {
			break
		}

		if m == l || m == r {
			if target > nums[r] {
				return r + 1
			} else if target > nums[l] {
				return r
			} else {
				return l
			}
		}
	}

	return m
}
