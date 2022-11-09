/**
 * Author: Deean
 * Date: 2022-11-09 23:26
 * FileName: algorithm/P1876. 长度为三且各字符不同的子字符串.go
 * Description:
 */

package main

import "fmt"

func countGoodSubstrings(s string) int {
	cnt := 0
	for i := 2; i < len(s); i++ {
		if s[i-2] != s[i-1] && s[i-2] != s[i] && s[i-1] != s[i] {
			cnt += 1
		}
	}
	return cnt
}

func main() {
	fmt.Println(countGoodSubstrings("aababcabc"))
}
