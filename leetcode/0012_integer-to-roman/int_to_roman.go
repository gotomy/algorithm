package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(10))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(1994))
}

func intToRoman(num int) string {
	romanMap := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	specialMap := map[int]string{
		4:   "IV",
		9:   "IX",
		40:  "XL",
		90:  "XC",
		400: "CD",
		900: "CM",
	}
	if val, ok := specialMap[num]; ok {
		return val
	}

	letterNums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	ret := make([]string, 0)
	for num != 0 {
		for _, letterNum := range letterNums {
			for i := 0; i < num/letterNum; i++ {
				ret = append(ret, romanMap[letterNum])
			}
			num = num % letterNum
		}
	}

	return strings.Join(ret, "")
}
