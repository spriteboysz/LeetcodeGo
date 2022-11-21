/**
 * Author: Deean
 * Date: 2022/11/21 23:23
 * FileName: algorithm/剑指 Offer II 018. 有效的回文.go
 * Description:
 */

package main

import (
	"fmt"
	"unicode"
)

func isPalindrome2(s string) bool {
	ss := []rune{}
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			ss = append(ss, unicode.ToLower(c))
		}
	}
	for i := 0; i < len(ss)/2; i++ {
		if ss[i] != ss[len(ss)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome2("A man, a plan, a canal: Panama"))
}
