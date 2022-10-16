/**
 * Author: Deean
 * Date: 2022-10-16 21:01
 * FileName: algorithm/P0338. 比特位计数.go
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
