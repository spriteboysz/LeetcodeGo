/**
 * Author: Deean
 * Date: 2022-10-20 23:32
 * FileName: algorithm/P0852. 山脉数组的峰顶索引.go
 * Description:
 */

package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	for i := 1; i < len(arr)-1; i++ {
		if arr[i-1] < arr[i] && arr[i+1] < arr[i] {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{3, 4, 5, 1}
	fmt.Println(peakIndexInMountainArray(arr))
}
