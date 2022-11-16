/**
 * Author: Deean
 * Date: 2022/11/16 22:03
 * FileName: interview/面试题 17.04. 消失的数字.go
 * Description:
 */

package main

import "fmt"

func missingNumber(nums []int) int {
	sum, n := 0, len(nums)
	for _, num := range nums {
		sum += num
	}
	return n*(n+1)/2 - sum
}

func main() {
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
