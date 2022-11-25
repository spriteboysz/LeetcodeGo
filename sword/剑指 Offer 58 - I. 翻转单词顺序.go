/**
 * Author: Deean
 * Date: 2022/11/25 22:25
 * FileName: sword/剑指 Offer 58 - I. 翻转单词顺序.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	n := len(words)
	for i := 0; i < n/2; i++ {
		words[i], words[n-1-i] = words[n-1-i], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseWords("a good   example"))
}
