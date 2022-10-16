/**
 * Author: Deean
 * Date: 2022-10-16 23:34
 * FileName: algorithm/P1374. 生成每种字符都是奇数个的字符串.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func generateTheString(n int) string {
	if n%2 == 0 {
		return strings.Repeat("a", n-1) + "b"
	} else {
		return strings.Repeat("a", n)
	}
}

func main() {
	fmt.Println(generateTheString(2))
	fmt.Println(generateTheString(7))
	fmt.Println(generateTheString(1))
}
