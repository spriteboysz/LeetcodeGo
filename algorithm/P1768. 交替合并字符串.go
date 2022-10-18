/**
 * Author: Deean
 * Date: 2022-10-18 23:23
 * FileName: algorithm/P1768. 交替合并字符串.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	var ss strings.Builder
	n1, n2 := len(word1), len(word2)
	for i := 0; i < n1 || i < n2; i++ {
		if i < n1 {
			ss.WriteByte(word1[i])
		}
		if i < n2 {
			ss.WriteByte(word2[i])
		}
	}
	return ss.String()
}

func main() {
	fmt.Println(mergeAlternately("ab", "pqrs"))
	fmt.Println(mergeAlternately("abcd", "pq"))
}
