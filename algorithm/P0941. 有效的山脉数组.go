/**
 * Author: Deean
 * Date: 2022/11/24 23:44
 * FileName: algorithm/P0941. 有效的山脉数组.go
 * Description:
 */

package main

import "fmt"

func validMountainArray(arr []int) bool {
	left, right := 0, len(arr)-1
	for left+1 < len(arr) && arr[left] < arr[left+1] {
		left++
	}
	for right > 0 && arr[right-1] > arr[right] {
		right--
	}
	if left > 0 && right < len(arr)-1 && left == right {
		return true
	}
	return false
}

func main() {
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))
}
