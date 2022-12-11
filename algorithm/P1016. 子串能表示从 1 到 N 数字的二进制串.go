/**
 * Author: Deean
 * Date: 2022/12/11 17:39
 * FileName: algorithm/P1016. 子串能表示从 1 到 N 数字的二进制串.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func queryString(s string, n int) bool {
	for i := 1; i <= n; i++ {
		cur := fmt.Sprintf("%b", i)
		if !strings.Contains(s, cur) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(queryString("0110", 3))
	fmt.Println(queryString("0110", 4))
}
