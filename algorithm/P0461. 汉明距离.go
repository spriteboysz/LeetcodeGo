/**
 * Author: Deean
 * Date: 2022-10-16 10:33
 * FileName: algorithm/P0461. 汉明距离.go
 * Description:
 */

package main

import (
	"fmt"
	"math/bits"
)

func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

func main() {
	fmt.Println(hammingDistance(3, 1))
}
