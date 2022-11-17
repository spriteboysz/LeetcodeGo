/**
 * Author: Deean
 * Date: 2022/11/17 23:17
 * FileName: algorithm/P0557. 反转字符串中的单词 III.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		reverse := []string{}
		for _, c := range word {
			reverse = append([]string{string(c)}, reverse...)
		}
		words[i] = strings.Join(reverse, "")
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}
