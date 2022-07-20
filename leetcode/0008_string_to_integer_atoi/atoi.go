package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(myAtoi("20000000000000000000"))
}

func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	// 空格 32
	// 48-57分别是0-9
	// + 43
	// - 45
	//
	sign := 1
	// 第一次出现非零数字标志
	firstNum := false
	fhFlag := false

	nums := make([]rune, 0)

	for idx, ch := range s {
		if idx == 0 {
			if ch > 57 || (ch < 48 && ch != 43 && ch != 45) {
				return 0
			}
		}
		if fhFlag {
			if ch > 57 || ch < 48 {
				return 0
			}
			fhFlag = false
		}
		if firstNum {
			if ch < 48 || ch > 57 {
				break
			} else {
				nums = append(nums, ch)
			}
		} else {
			if ch == 43 || ch == 45 {
				if ch == 45 {
					sign = -1
				}
				fhFlag = true
			} else {
				if ch >= 48 && ch <= 57 {
					if !firstNum {
						firstNum = true
					}
					nums = append(nums, ch)
				}
			}
		}

	}

	maxInt := math.Pow(2, 31) - 1
	minInt := math.Pow(-2, 31)

	ret := 0
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		tmpInt := (int(nums[i]) - 48)
		if tmpInt != 0 {
			if (len(nums) - i - 1) > 10 {
				return int(maxInt)
			}
		}
		ret = ret + tmpInt*(int(math.Pow(10, float64(len(nums)-i-1))))
		fmt.Println(ret)
		if sign == 1 {
			if ret > int(maxInt) {
				return int(maxInt)
			}
		} else {
			if ret*-1 < int(minInt) {
				return int(minInt)
			}
		}
	}

	return ret * sign
}
