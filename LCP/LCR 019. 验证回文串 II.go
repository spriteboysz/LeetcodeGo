/**
 * Author: Deean
 * Date: 2023-10-14 22:21
 * FileName: LCP/LCR 019. 验证回文串 II.go
 * Description:
 */

package main

import "fmt"

func validPalindrome(s string) bool {
	var check = func(s string) bool {
		for i, n := 0, len(s); i < n/2; i++ {
			if s[i] != s[n-1-i] {
				return false
			}
		}
		return true
	}

	for i, n := 0, len(s); i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return check(s[0:i]+s[i+1:]) || check(s[0:n-1-i]+s[n-i:])
		}
	}
	return true
}

func main() {
	fmt.Println(validPalindrome("abca"))
}
