/**
 * Author: Deean
 * Date: 2022/12/11 16:55
 * FileName: algorithm/P1324. 竖直打印单词.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func printVertically(s string) []string {
	words := strings.Split(s, " ")
	chars := make([]byte, len(words))
	blank := strings.Repeat(" ", len(words))
	vertical := make([]string, 0)
	for true {
		for i, word := range words {
			if len(vertical) < len(word) {
				chars[i] = word[len(vertical)]
			} else {
				chars[i] = ' '
			}
		}
		if string(chars) == blank {
			return vertical
		}
		vertical = append(vertical, strings.TrimRight(string(chars), " "))
	}
	return vertical
}

func main() {
	fmt.Println(printVertically("TO BE OR NOT TO BE"))
}
