/**
 * Author: Deean
 * Date: 2022/11/23 22:36
 * FileName: algorithm/P0058. 最后一个单词的长度.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	arr := strings.Split(s, " ")
	return len(arr[len(arr)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
}
