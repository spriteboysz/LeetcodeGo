/**
 * Author: Deean
 * Date: 2022/12/4 17:19
 * FileName: algorithm/P2138. 将字符串拆分为若干长度为 k 的组.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func divideString(s string, k int, fill byte) []string {
	n, divided := len(s), []string{}
	for i := 0; i < n; i += k {
		if i+k <= n {
			divided = append(divided, s[i:i+k])
		} else {
			divided = append(divided, s[i:n]+strings.Repeat(string(fill), k-(n-i)))
		}
	}
	return divided
}

func main() {
	fmt.Println(divideString("abcdefghij", 3, 'x'))
}
