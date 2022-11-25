/**
 * Author: Deean
 * Date: 2022/11/25 23:20
 * FileName: interview/面试题 01.09. 字符串轮转.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func isFlipedString(s1 string, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s1+s1, s2)
}

func main() {
	fmt.Println(isFlipedString("waterbottle", "erbottlewat"))
}
