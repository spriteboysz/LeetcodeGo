/**
 * Author: Deean
 * Date: 2022-10-16 20:08
 * FileName: algorithm/P1967. 作为子字符串出现在单词中的字符串数目.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func numOfStrings(patterns []string, word string) int {
	cnt := 0
	for _, pattern := range patterns {
		if strings.Contains(word, pattern) {
			cnt++
		}
	}
	return cnt
}

func main() {
	patterns := []string{"a", "b", "c"}
	fmt.Println(numOfStrings(patterns, "aabb"))
}
