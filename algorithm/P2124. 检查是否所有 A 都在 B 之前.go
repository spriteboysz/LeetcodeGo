/**
 * Author: Deean
 * Date: 2022-10-20 23:58
 * FileName: algorithm/P2124. 检查是否所有 A 都在 B 之前.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func checkString(s string) bool {
	if strings.Contains(s, "ba") {
		return false
	}
	return true
}

func main() {
	fmt.Println(checkString("abab"))
	fmt.Println(checkString("bbb"))
}
