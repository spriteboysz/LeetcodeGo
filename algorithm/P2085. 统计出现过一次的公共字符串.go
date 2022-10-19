/**
 * Author: Deean
 * Date: 2022-10-19 23:21
 * FileName: algorithm/P2085. 统计出现过一次的公共字符串.go
 * Description:
 */

package main

import "fmt"

func countWords(words1 []string, words2 []string) int {
	hash1, hash2 := map[string]int{}, map[string]int{}
	for _, word := range words1 {
		hash1[word]++
	}
	for _, word := range words2 {
		hash2[word]++
	}
	cnt := 0
	for k, v := range hash1 {
		if v == 1 && hash2[k] == 1 {
			cnt++
		}
	}
	return cnt
}

func main() {
	words1 := []string{"a", "ab"}
	words2 := []string{"a", "a", "a", "ab"}
	fmt.Println(countWords(words1, words2))
}
