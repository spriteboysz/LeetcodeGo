/**
 * Author: Deean
 * Date: 2022/11/26 23:06
 * FileName: algorithm/P1544. 整理字符串.go
 * Description:
 */

package main

import (
	"fmt"
)

func makeGood(s string) string {
	check := func(a, b byte) bool {
		if a == b {
			return false
		}
		if a > b {
			return b+32 == a
		}
		return a+32 == b
	}
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if len(stack) != 0 && check(stack[len(stack)-1], s[i]) {
			stack = stack[0 : len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	res := ""
	for i := 0; i < len(stack); i++ {
		res += string(stack[i])
	}
	return res
}

func main() {
	fmt.Println(makeGood("abBAcC"))
}
