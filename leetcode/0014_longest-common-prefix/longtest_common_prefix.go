package main

import "fmt"

func main() {
	strs := []string{"a"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	ret := ""
	strMap := make(map[int]string)
	for idx, val := range strs {
		if val == "" {
			return ""
		}
		strMap[idx] = val
	}
	idx := 0
	flag := true
	temp := ""
	for flag {
		tempMap := make(map[string]int)
		for i := 0; i < len(strs); i++ {
			tempStr := strMap[i]

			if len(tempStr) <= idx {
				flag = false
				break
			}
			temp = tempStr[0 : idx+1]

			if i == 0 {
				tempMap[temp] = 1
			}

			if _, ok := tempMap[temp]; ok {

			} else {
				flag = false
				break
			}
		}
		if flag {
			ret = temp
		}
		idx++
	}
	return ret
}
