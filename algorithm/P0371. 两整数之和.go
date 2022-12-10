/**
 * Author: Deean
 * Date: 2022/12/10 16:51
 * FileName: algorithm/P0371. 两整数之和.go
 * Description:
 */

package main

import "fmt"

func getSum(a int, b int) int {
	for b != 0 {
		carry := uint(a&b) << 1
		a ^= b
		b = int(carry)
	}
	return a
}

func main() {
	fmt.Println(getSum(2, 3))
}
