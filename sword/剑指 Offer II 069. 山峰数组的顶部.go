/**
 * Author: Deean
 * Date: 2022-10-20 23:17
 * FileName: sword/剑指 Offer II 069. 山峰数组的顶部.go
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
	arr := []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}
	fmt.Println(peakIndexInMountainArray(arr))
}
