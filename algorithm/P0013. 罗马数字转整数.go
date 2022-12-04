/**
 * Author: Deean
 * Date: 2022/12/4 17:59
 * FileName: algorithm/P0013. 罗马数字转整数.go
 * Description:
 */

package main

import "fmt"

func romanToInt(s string) int {
	var chars2 = []string{"I", "II", "III", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	var nums2 = []int{1, 2, 3, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	num, index := 0, len(chars2)-1
	for s != "" {
		temp := chars2[index]
		n := len(temp)
		if n > len(s) {
			index--
			continue
		}
		if s[:n] == temp {
			s = s[n:]
			num += nums2[index]
		} else {
			index--
		}
	}
	return num
}

func main() {
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
