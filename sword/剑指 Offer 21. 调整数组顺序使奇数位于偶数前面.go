/**
 * Author: Deean
 * Date: 2022/11/15 23:01
 * FileName: sword/剑指 Offer 21. 调整数组顺序使奇数位于偶数前面.go
 * Description:
 */

package main

import "fmt"

func exchange(nums []int) []int {
	even, odd := []int{}, []int{}
	for _, num := range nums {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	return append(odd, even...)
}

func main() {
	fmt.Println(exchange([]int{1, 2, 3, 4}))
}
