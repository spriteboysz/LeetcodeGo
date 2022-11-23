/**
 * Author: Deean
 * Date: 2022/11/22 23:26
 * FileName: algorithm/P1592. 重新排列单词间的空格.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func reorderSpaces(text string) string {
	words := strings.Fields(text)
	cnt := len(text)
	for _, word := range words {
		cnt -= len(word)
	}
	if len(words)-1 == 0 {
		return words[0] + strings.Repeat(" ", cnt)
	}
	div, mod := cnt/(len(words)-1), cnt%(len(words)-1)
	separator, suffix := strings.Repeat(" ", div), strings.Repeat(" ", mod)
	return strings.Join(words, separator) + suffix
}

func main() {
	fmt.Println(reorderSpaces("  this   is  a sentence "))
	fmt.Println(reorderSpaces(" practice   makes   perfect"))
}
