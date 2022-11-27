/**
 * Author: Deean
 * Date: 2022/11/27 22:05
 * FileName: algorithm/P0009. 回文数.go
 * Description:
 */

package main

import "fmt"

func isPalindromeNum(x int) bool {
	if x < 0 {
		return false
	}
	x1, x2 := x, 0
	for x1 > 0 {
		mod := x1 % 10
		x2 = 10*x2 + mod
		x1 /= 10
	}
	return x == x2
}

func main() {
	fmt.Println(isPalindromeNum(121))
	fmt.Println(isPalindromeNum(-121))
	fmt.Println(isPalindromeNum(10))
}
