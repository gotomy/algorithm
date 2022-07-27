package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	specialMap := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	sLen := len(s)
	ret := 0
	i := 0
	for i < sLen {
		str := s[i : i+1]
		if str == "I" || str == "X" || str == "C" {
			if i+2 < sLen+1 {
				str = s[i : i+2]
				if val, ok := specialMap[str]; ok {
					ret += val
					i = i + 2
				} else {
					str = s[i : i+1]
					val, _ := romanMap[str]
					ret += val
					i = i + 1
				}
			} else {
				val, _ := romanMap[str]
				ret += val
				i = i + 1
			}
		} else {
			val, _ := romanMap[str]
			ret += val
			i = i + 1
		}

	}
	return ret
}
