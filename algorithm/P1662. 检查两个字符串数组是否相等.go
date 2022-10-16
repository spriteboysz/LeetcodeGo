/**
 * Author: Deean
 * Date: 2022-10-16 14:25
 * FileName: algorithm/P1662. 检查两个字符串数组是否相等.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}

func main() {
	word1 := []string{"abc", "d", "defg"}
	word2 := []string{"abcddefg"}
	fmt.Println(arrayStringsAreEqual(word1, word2))
}
