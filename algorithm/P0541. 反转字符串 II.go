/**
 * Author: Deean
 * Date: 2022/11/19 19:58
 * FileName: algorithm/P0541. 反转字符串 II.go
 * Description:
 */

package main

import "fmt"

func reverseStr(s string, k int) string {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	ss := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		sub := ss[i:min(i+k, len(s))]
		for j, n := 0, len(sub); j < n/2; j++ {
			sub[j], sub[n-1-j] = sub[n-1-j], sub[j]
		}
	}
	return string(ss)
}

func main() {
	fmt.Println(reverseStr("abcdefg", 2))
	fmt.Println(reverseStr("abcd", 2))
}
