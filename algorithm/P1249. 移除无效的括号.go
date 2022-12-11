/**
 * Author: Deean
 * Date: 2022/12/11 16:23
 * FileName: algorithm/P1249. 移除无效的括号.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func minRemoveToMakeValid(s string) string {
	n, sum, ss := len(s), 0, []byte(s)
	for i := 0; i < n; i++ {
		if ss[i] == '(' {
			sum++
		} else if ss[i] == ')' {
			sum--
			if sum < 0 {
				ss[i] = '0'
				sum = 0
			}
		}
	}
	sum = 0
	for i := n - 1; i >= 0; i-- {
		if ss[i] == ')' {
			sum++
		} else if ss[i] == '(' {
			sum--
			if sum < 0 {
				ss[i] = '0'
				sum = 0
			}
		}
	}
	return strings.ReplaceAll(string(ss), "0", "")
}

func main() {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)"))
}
