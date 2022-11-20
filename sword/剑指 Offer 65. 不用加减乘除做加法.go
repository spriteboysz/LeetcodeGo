/**
 * Author: Deean
 * Date: 2022/11/20 22:21
 * FileName: sword/剑指 Offer 65. 不用加减乘除做加法.go
 * Description:
 */

package main

import "fmt"

func add(a int, b int) int {
	for b != 0 {
		carry := uint(a&b) << 1
		a ^= b
		b = int(carry)
	}
	return a
}

func main() {
	fmt.Println(add(1, 1))
}
