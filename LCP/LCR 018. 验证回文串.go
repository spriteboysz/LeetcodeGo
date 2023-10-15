/**
 * Author: Deean
 * Date: 2023-10-14 22:16
 * FileName: LCP/LCR 018. 验证回文串.go
 * Description:
 */

package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	ss := []rune{}
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			ss = append(ss, unicode.ToLower(c))
		}
	}
	for i, n := 0, len(ss); i < n/2; i++ {
		if ss[i] != ss[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
