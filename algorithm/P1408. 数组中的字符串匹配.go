/**
 * Author: Deean
 * Date: 2022-10-21 23:36
 * FileName: algorithm/P1408. 数组中的字符串匹配.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func stringMatching(words []string) []string {
	text := strings.Join(words, "#")
	target := []string{}
	for _, word := range words {
		if strings.Index(text, word) != strings.LastIndex(text, word) {
			target = append(target, word)
		}
	}
	return target
}

func main() {
	words := []string{"leetcode", "et", "code"}
	fmt.Println(stringMatching(words))
}
