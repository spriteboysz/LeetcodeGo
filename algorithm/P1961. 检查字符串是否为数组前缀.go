/**
 * Author: Deean
 * Date: 2022/12/14 23:39
 * FileName: algorithm/P1961. 检查字符串是否为数组前缀.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func isPrefixString(s string, words []string) bool {
	strs, sum := []string{}, 0
	for _, word := range words {
		strs = append(strs, word)
		sum += len(word)
		if sum >= len(s) {
			break
		}
	}
	ss := strings.Join(strs, "")
	return strings.EqualFold(s, ss)
}

func main() {
	fmt.Println(isPrefixString("iloveleetcode", []string{"i", "love", "leetcode", "apples"}))
}
