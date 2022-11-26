/**
 * Author: Deean
 * Date: 2022/11/26 22:20
 * FileName: algorithm/P0680. 验证回文串 II.go
 * Description:
 */

package main

import "fmt"

func validPalindrome(s string) bool {
	check := func(left, right int) bool {
		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}

	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		if s[left] != s[right] {
			return check(left+1, right) || check(left, right-1)
		}
	}
	return true
}

func main() {
	fmt.Println(validPalindrome("abca"))
}
