/**
 * Author: Deean
 * Date: 2022/12/4 16:24
 * FileName: algorithm/P1763. 最长的美好子字符串.go
 * Description:
 */

package main

import (
	"fmt"
	"unicode"
)

func longestNiceSubstring(s string) string {
	nice := ""
	for i := range s {
		lower, upper := 0, 0
		for j := i; j < len(s); j++ {
			if unicode.IsLower(rune(s[j])) {
				lower |= 1 << (s[j] - 'a')
			} else {
				upper |= 1 << (s[j] - 'A')
			}
			if lower == upper && j-i+1 > len(nice) {
				nice = s[i : j+1]
			}
		}
	}
	return nice
}

func main() {
	fmt.Println(longestNiceSubstring("YazaAay"))
}
