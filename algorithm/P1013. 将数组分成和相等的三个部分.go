/**
 * Author: Deean
 * Date: 2022/11/24 23:34
 * FileName: algorithm/P1013. 将数组分成和相等的三个部分.go
 * Description:
 */

package main

import "fmt"

func canThreePartsEqualSum(arr []int) bool {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	if sum%3 != 0 {
		return false
	}
	target := sum / 3
	n, i, cur, count := len(arr), 0, 0, 0
	for i < n {
		cur += arr[i]
		if cur == target {
			cur = 0
			count++
		}
		i++
	}
	if count < 3 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canThreePartsEqualSum([]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}))
}
