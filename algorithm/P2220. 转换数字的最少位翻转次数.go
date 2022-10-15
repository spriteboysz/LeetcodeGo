/**
 * Author: Deean
 * Date: 2022-10-15 22:05
 * FileName: algorithm/P2220. 转换数字的最少位翻转次数.go
 * Description:
 */

package main

import (
	"fmt"
	"math/bits"
)

func minBitFlips(start int, goal int) int {
	return bits.OnesCount(uint(start ^ goal))
}

func main() {
	fmt.Println(minBitFlips(3, 4))
}
