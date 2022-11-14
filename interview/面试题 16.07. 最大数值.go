/**
 * Author: Deean
 * Date: 2022/11/14 23:25
 * FileName: interview/面试题 16.07. 最大数值.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
)

func maximumII(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func main() {
	fmt.Println(maximumII(3, 2))
}
