/**
 * Author: Deean
 * Date: 2022-10-16 22:05
 * FileName: algorithm/P2108. 找出数组中的第一个回文字符串.go
 * Description:
 */

package main

import "fmt"

func firstPalindrome(words []string) string {
	isPalindrome := func(word string) bool {
		left, right := 0, len(word)-1
		for left < right {
			if word[left] != word[right] {
				return false
			}
			left++
			right--
		}
		return true
	}
	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}
	return ""
}

func main() {
	words := []string{"notapalindrome", "racecar"}
	fmt.Println(firstPalindrome(words))
}
