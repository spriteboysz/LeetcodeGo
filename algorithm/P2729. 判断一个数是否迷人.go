/**
 * Author: Deean
 * Date: 2023-06-24 15:58
 * FileName: algorithm/P2729. 判断一个数是否迷人.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
)

func isFascinating(n int) bool {
	ss := strconv.Itoa(n) + strconv.Itoa(n*2) + strconv.Itoa(n*3)
	if len(ss) != 9 {
		return false
	}
	digits := make([]int, 10)
	for _, c := range ss {
		digits[c-'0'] += 1
	}
	for i := 1; i < 10; i++ {
		if digits[i] != 1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isFascinating(192))
	fmt.Println(isFascinating(100))
}
