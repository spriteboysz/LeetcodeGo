/**
 * Author: Deean
 * Date: 2022-11-09 22:39
 * FileName: algorithm/P1935. 可以输入的最大单词数.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func canBeTypedWords(text string, brokenLetters string) int {
	words := strings.Split(text, " ")
	cnt := len(words)
	for _, word := range words {
		for _, c := range brokenLetters {
			if strings.Contains(word, string(c)) {
				cnt -= 1
				break
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(canBeTypedWords("leet code", "lt"))
	fmt.Println(canBeTypedWords("leet code", "e"))
}
