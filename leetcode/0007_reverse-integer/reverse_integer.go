package main

import "fmt"

func main() {
	fmt.Println(reverse(3903949))
}

func reverse(x int) int {
	ret := 0
	for x != 0 {
		temp := x % 10
		if ret > 214748364 || (ret == 214748364 && temp > 7) {
			return 0
		}
		if ret < -214748364 || (ret == -214748364 && temp < -8) {
			return 0
		}
		ret = ret*10 + temp
		x = x / 10
	}
	return ret
}
