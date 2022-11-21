/**
 * Author: Deean
 * Date: 2022/11/21 22:53
 * FileName: algorithm/P1796. 字符串中第二大的数字.go
 * Description:
 */

package main

import (
	"fmt"
	"unicode"
)

func secondHighest(s string) int {
	number := [10]bool{}
	for _, c := range s {
		if unicode.IsDigit(c) {
			number[c-'0'] = true
		}
	}
	flag := false
	for i := 9; i >= 0; i-- {
		if flag && number[i] {
			return i
		}
		if number[i] {
			flag = true
		}
	}
	return -1
}

func main() {
	fmt.Println(secondHighest("dfa12321afd"))
}
