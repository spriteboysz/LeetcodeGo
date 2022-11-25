/**
 * Author: Deean
 * Date: 2022/11/25 22:48
 * FileName: interview/面试题 05.06. 整数转换.go
 * Description:
 */

package main

import (
	"fmt"
)

func convertInteger(A int, B int) int {
	cnt := 0
	for i := 0; i < 32; i++ {
		cnt += (A & 1) ^ (B & 1)
		A >>= 1
		B >>= 1
	}
	return cnt
}

func main() {
	fmt.Println(convertInteger(29, 15))
}
