/**
 * Author: Deean
 * Date: 2022/11/24 23:22
 * FileName: algorithm/P0345. 反转字符串中的元音字母.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	vowels := []rune{}
	for _, c := range s {
		if strings.Contains("aeiouAEIOU", string(c)) {
			vowels = append(vowels, c)
		}
	}
	reversed, index := []string{}, len(vowels)-1
	for _, c := range s {
		if strings.Contains("aeiou", string(c)) {
			reversed = append(reversed, string(vowels[index]))
			index--
		} else {
			reversed = append(reversed, string(c))
		}
	}
	return strings.Join(reversed, "")
}

func main() {
	fmt.Println(reverseVowels("leetcode"))
}
