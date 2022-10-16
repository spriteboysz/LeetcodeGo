/**
 * Author: Deean
 * Date: 2022-10-16 17:50
 * FileName: algorithm/P2255. 统计是给定字符串前缀的字符串数目.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func countPrefixes(words []string, s string) int {
	cnt := 0
	for _, word := range words {
		if strings.HasPrefix(s, word) {
			cnt++
		}
	}
	return cnt
}

func main() {
	words := []string{"a", "a"}
	fmt.Println(countPrefixes(words, "aa"))
}
