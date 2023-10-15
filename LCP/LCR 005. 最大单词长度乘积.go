/**
 * Author: Deean
 * Date: 2023-10-14 21:38
 * FileName: LCP/LCR 005. 最大单词长度乘积.go
 * Description:
 */

package main

import "fmt"

func maxProduct(words []string) int {
	hash := map[string]int{}
	for _, word := range words {
		v := 0
		for _, c := range word {
			v |= 1 << (c - 'a')
		}
		hash[word] = v
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
	fmt.Println(maxProduct([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}))
}
