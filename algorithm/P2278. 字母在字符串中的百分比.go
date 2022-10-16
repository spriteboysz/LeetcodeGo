/**
 * Author: Deean
 * Date: 2022-10-16 14:21
 * FileName: algorithm/P2278. 字母在字符串中的百分比.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func percentageLetter(s string, letter byte) int {
	return strings.Count(s, string(letter)) * 100 / len(s)
}

func main() {
	fmt.Println(percentageLetter("foobar", 'o'))
}
