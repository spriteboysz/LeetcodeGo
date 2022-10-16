/**
 * Author: Deean
 * Date: 2022-10-16 20:04
 * FileName: algorithm/P2185. 统计包含给定前缀的字符串.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func prefixCount(words []string, pref string) int {
	cnt := 0
	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			cnt++
		}
	}
	return cnt
}

func main() {
	words := []string{"leetcode", "win", "loops", "success"}
	fmt.Println(prefixCount(words, "code"))
}
