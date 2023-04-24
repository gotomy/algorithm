package main

import "fmt"

func main() {
	nums := []int{2}
	fmt.Println(removeElement(nums, 3))
}

func removeElement(nums []int, val int) int {
	k := len(nums) - 1
	for k >= 0 {
		j := k - 1
		v := nums[k]
		if v == val {
			nums[k] = -1
		} else {
			for j >= 0 {
				s := nums[j]
				if s == val {
					nums[j] = nums[k]
					nums[k] = -1
					break
				}
				j--
			}
		}
		k--
	}
	fmt.Println(nums)
	ret := 0
	for i, v := range nums {
		if v == -1 {
			ret = i
			return ret
		}
	}

	if ret == 0 {
		return len(nums)
	}

	return 0
}
