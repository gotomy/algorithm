package main

import "fmt"

func main() {
	nums := []int{1}
	ret := searchRange(nums, 1)
	fmt.Println(ret)
}

func searchRange(nums []int, target int) []int {
	ret := []int{-1, -1}
	k := 0
	j := len(nums) - 1
	fisrstMeet := false
	lastMeet := false
	for k < j+1 {
		if nums[k] != target {
			k++
		} else {
			ret[0] = k
			fisrstMeet = true
		}

		if nums[j] != target {
			j--
		} else {
			ret[1] = j
			lastMeet = true
		}

		if fisrstMeet && lastMeet {
			break
		}
	}

	return ret
}
