/**
 * Author: Deean
 * Date: 2022/11/16 23:10
 * FileName: algorithm/P2243. 计算字符串的数字和.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func digitSum(s string, k int) string {
	for len(s) > k {
		var build strings.Builder
		n := len(s)
		for i := 0; i < n; i += k {
			sum := 0
			for j := 0; j < k && i+j < n; j++ {
				sum += int(s[i+j] - '0')
			}
			build.WriteString(strconv.Itoa(sum))
		}
		s = build.String()
	}
	return s
}

func main() {
	fmt.Println(digitSum("11111222223", 3))
}
