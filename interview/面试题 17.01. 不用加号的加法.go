/**
 * Author: Deean
 * Date: 2022/11/20 22:18
 * FileName: interview/面试题 17.01. 不用加号的加法.go
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
