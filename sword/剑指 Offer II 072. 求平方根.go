/**
 * Author: Deean
 * Date: 2022/11/24 23:30
 * FileName: sword/剑指 Offer II 072. 求平方根.go
 * Description:
 */

package main

import "fmt"

func mySqrt(x int) int {
	left, right, sqrt := 0, x, -1
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			sqrt = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return sqrt
}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
}
