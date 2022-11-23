/**
 * Author: Deean
 * Date: 2022/11/23 22:22
 * FileName: algorithm/P0434. 字符串中的单词数.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func countSegments(s string) int {
	if s == "" {
		return 0
	}
	return len(strings.Fields(s))
}

func main() {
	fmt.Println(countSegments("Hello, my name is John"))
	fmt.Println(countSegments("  , , , ,        a, eaefa "))
	fmt.Println(countSegments(""))
}
