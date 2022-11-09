/**
 * Author: Deean
 * Date: 2022-11-09 22:01
 * FileName: algorithm/P0884. 两句话中的不常见单词.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func uncommonFromSentences(s1 string, s2 string) []string {
	hash1, hash2 := map[string]int{}, map[string]int{}
	for _, word := range strings.Split(s1, " ") {
		hash1[word] += 1
	}
	for _, word := range strings.Split(s2, " ") {
		hash2[word] += 1
	}
	uncommon := []string{}
	for word := range hash1 {
		if hash1[word] == 1 && hash2[word] == 0 {
			uncommon = append(uncommon, word)
		}
	}
	for word := range hash2 {
		if hash2[word] == 1 && hash1[word] == 0 {
			uncommon = append(uncommon, word)
		}
	}
	return uncommon
}

func main() {
	fmt.Println(uncommonFromSentences("apple apple", "banana"))
}
