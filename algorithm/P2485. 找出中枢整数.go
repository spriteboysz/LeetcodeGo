/**
 * Author: Deean
 * Date: 2022/12/1 22:44
 * FileName: algorithm/P2485. 找出中枢整数.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
)

func pivotInteger(n int) int {
	m := n * (n + 1) / 2
	x := int(math.Sqrt(float64(m)))
	if x*x == m {
		return x
	}
	return -1
}

func main() {
	fmt.Println(pivotInteger(8))
}
