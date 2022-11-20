/**
 * Author: Deean
 * Date: 2022/11/20 17:21
 * FileName: algorithm/P2027. 转换字符串的最少操作次数.go
 * Description:
 */

package main

import "fmt"

func minimumMoves(s string) int {
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'X' {
			cnt++
			i += 2
		}
	}
	return cnt
}

func main() {
	fmt.Println(minimumMoves("XXOX"))
}
