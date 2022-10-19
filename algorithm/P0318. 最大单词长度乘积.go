/**
 * Author: Deean
 * Date: 2022-10-19 00:04
 * FileName: algorithm/P0318. 最大单词长度乘积.go
 * Description:
 */

package main

import "fmt"

func maxLenProduct(words []string) int {
	hash := map[string]int{}
	for _, word := range words {
		value := 0
		for _, c := range word {
			value |= 1 << (c - 'a')
		}
		hash[word] = value
	}
	maximum := 0
	for i, word := range words {
		for j := i + 1; j < len(words); j++ {
			if hash[word]&hash[words[j]] == 0 {
				if maximum < len(word)*len(words[j]) {
					maximum = len(word) * len(words[j])
				}
			}
		}
	}
	return maximum
}

func main() {
	words := []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}
	fmt.Println(maxLenProduct(words))
}
