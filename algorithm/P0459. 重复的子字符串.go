/**
 * Author: Deean
 * Date: 2022/11/21 23:08
 * FileName: algorithm/P0459. 重复的子字符串.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	return strings.Contains((s + s)[1:2*len(s)-1], s)
}

func main() {
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
}
