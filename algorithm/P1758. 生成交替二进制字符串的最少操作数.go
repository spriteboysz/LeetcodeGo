/**
 * Author: Deean
 * Date: 2022/11/16 23:04
 * FileName: algorithm/P1758. 生成交替二进制字符串的最少操作数.go
 * Description:
 */

package main

import "fmt"

func minOperations3(s string) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	s0, s1 := 0, 0
	for i := 0; i < len(s); i += 2 {
		if s[i] == '1' {
			s0++
		} else {
			s1++
		}
	}
	for i := 1; i < len(s); i += 2 {
		if s[i] == '1' {
			s1++
		} else {
			s0++
		}
	}
	return min(s1, s0)
}

func main() {
	fmt.Println(minOperations3("1111"))
}
