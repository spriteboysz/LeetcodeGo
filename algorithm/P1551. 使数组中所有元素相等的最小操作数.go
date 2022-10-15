/**
 * Author: Deean
 * Date: 2022-10-15 23:42
 * FileName: algorithm/P1551. 使数组中所有元素相等的最小操作数.go
 * Description:
 */

package main

import "fmt"

func minOperations(n int) int {
	abs := func(num int) int {
		if num >= 0 {
			return num
		}
		return -num
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += abs(n - (2*i + 1))
	}
	return sum / 2
}

func main() {
	fmt.Println(minOperations(6))
}
