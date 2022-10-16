/**
 * Author: Deean
 * Date: 2022-10-16 23:21
 * FileName: algorithm/P1332. 删除回文子序列.go
 * Description:
 */

package main

import "fmt"

func removePalindromeSub(s string) int {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return 2
		}
		left++
		right--
	}
	return 1
}

func main() {
	fmt.Println(removePalindromeSub("baabb"))
	fmt.Println(removePalindromeSub("ababa"))
}
