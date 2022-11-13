/**
 * Author: Deean
 * Date: 2022/11/13 22:40
 * FileName: algorithm/P1385. 两个数组间的距离值.go
 * Description:
 */

package main

import "fmt"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	abs := func(num int) int {
		if num >= 0 {
			return num
		}
		return -num
	}
	cnt := 0
	for _, num1 := range arr1 {
		flag := true
		for _, num2 := range arr2 {
			if abs(num1-num2) <= d {
				flag = false
				break
			}
		}
		if flag {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(findTheDistanceValue([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2))
}
