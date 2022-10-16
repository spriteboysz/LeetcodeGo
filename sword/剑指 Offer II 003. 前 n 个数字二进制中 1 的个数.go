/**
 * Author: Deean
 * Date: 2022-10-16 20:57
 * FileName: sword/剑指 Offer II 003. 前 n 个数字二进制中 1 的个数.go
 * Description:
 */

package main

import "fmt"

func countBits(n int) []int {
	var count []int
	for i := 0; i <= n; i++ {
		num, cur := i, 0
		for num > 0 {
			cur += num & 1
			num >>= 1
		}
		count = append(count, cur)
	}
	return count
}

func main() {
	fmt.Println(countBits(5))
}
