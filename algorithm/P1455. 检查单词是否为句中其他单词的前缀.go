/**
 * Author: Deean
 * Date: 2022/11/12 23:36
 * FileName: algorithm/P1455. 检查单词是否为句中其他单词的前缀.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	words := strings.Split(sentence, " ")
	for i, word := range words {
		if strings.HasPrefix(word, searchWord) {
			return i + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(isPrefixOfWord("this problem is an easy problem", "pro"))
}
