/**
 * Author: Deean
 * Date: 2022-10-19 23:46
 * FileName: sword/剑指 Offer II 005. 单词长度的最大乘积.go
 * Description:
 */

package main

import "fmt"

func maxProduct(words []string) int {
	hash := map[string]int{}
	for _, word := range words {
		value := 0
		for _, c := range word {
			value |= 1 << (c - 'a')
		}
		hash[word] = value
	}
	maximum := 0
	for i, word1 := range words {
		for _, word2 := range words[:i] {
			if hash[word1]&hash[word2] == 0 {
				if maximum < len(word1)*len(word2) {
					maximum = len(word1) * len(word2)
				}
			}
		}
	}
	return maximum
}

func main() {
	words := []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}
	fmt.Println(maxProduct(words))
}
