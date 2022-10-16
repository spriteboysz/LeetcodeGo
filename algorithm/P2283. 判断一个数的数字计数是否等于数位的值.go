/**
 * Author: Deean
 * Date: 2022-10-16 22:37
 * FileName: algorithm/P2283. 判断一个数的数字计数是否等于数位的值.go
 * Description:
 */

package main

import "fmt"

func digitCount(num string) bool {
	hash := map[rune]int{}
	for _, digit := range num {
		hash[digit-'0']++
	}
	for i, digit := range num {
		if hash[rune(i)] != int(digit-'0') {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(digitCount("1210"))
	fmt.Println(digitCount("030"))
}
