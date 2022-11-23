/**
 * Author: Deean
 * Date: 2022/11/23 22:10
 * FileName: algorithm/P0367. 有效的完全平方数.go
 * Description:
 */

package main

import "fmt"

func isPerfectSquare(num int) bool {
	left, right := 0, num
	for left <= right {
		mid := (left + right) / 2
		square := mid * mid
		if square < num {
			left = mid + 1
		} else if square > num {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14))
}
