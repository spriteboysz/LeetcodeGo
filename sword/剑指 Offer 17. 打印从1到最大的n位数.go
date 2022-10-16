/**
 * Author: Deean
 * Date: 2022-10-16 23:27
 * FileName: sword/剑指 Offer 17. 打印从1到最大的n位数.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
)

func printNumbers(n int) []int {
	var nums []int
	for i := 1; i < int(math.Pow10(n)); i++ {
		nums = append(nums, i)
	}
	return nums
}
func main() {
	fmt.Println(printNumbers(2))
}
