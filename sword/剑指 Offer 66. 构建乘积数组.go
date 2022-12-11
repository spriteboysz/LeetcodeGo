/**
 * Author: Deean
 * Date: 2022/12/11 16:09
 * FileName: sword/剑指 Offer 66. 构建乘积数组.go
 * Description:
 */

package main

import "fmt"

func constructArr(a []int) []int {
	n := len(a)
	if n == 0 {
		return []int{}
	}
	product := make([]int, n)
	product[0] = 1
	for i := 1; i < n; i++ {
		product[i] = a[i-1] * product[i-1]
	}
	for i, right := n-1, 1; i >= 0; i-- {
		product[i] *= right
		right *= a[i]
	}
	return product
}

func main() {
	fmt.Println(constructArr([]int{1, 2, 3, 4, 5}))
}
