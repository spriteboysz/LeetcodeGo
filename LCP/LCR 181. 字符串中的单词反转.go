/**
 * Author: Deean
 * Date: 2023-10-19 23:40
 * FileName: LCR/LCR 181. 字符串中的单词反转.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func reverseMessage(message string) string {
	words := strings.Fields(message)
	n := len(words)
	for i := 0; i < n/2; i++ {
		words[i], words[n-1-i] = words[n-1-i], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseMessage("the sky is blue"))
}
