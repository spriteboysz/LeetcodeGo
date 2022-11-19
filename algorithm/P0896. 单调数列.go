/**
 * Author: Deean
 * Date: 2022/11/19 22:21
 * FileName: algorithm/P0896. 单调数列.go
 * Description:
 */

package main

import "fmt"

func isMonotonic(nums []int) bool {
	inc, dec := true, true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			inc = false
		}
		if nums[i] < nums[i+1] {
			dec = false
		}
	}
	return inc || dec
}

func main() {
	fmt.Println(isMonotonic([]int{6, 5, 4, 4}))
}
