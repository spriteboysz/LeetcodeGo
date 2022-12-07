/**
 * Author: Deean
 * Date: 2022/12/6 23:21
 * FileName: algorithm/P2161. 根据给定数字划分数组.go
 * Description:
 */

package main

import "fmt"

func pivotArray(nums []int, pivot int) []int {
	lt, eq, gt := []int{}, []int{}, []int{}
	for _, num := range nums {
		if num < pivot {
			lt = append(lt, num)
		} else if num > pivot {
			gt = append(gt, num)
		} else {
			eq = append(eq, num)
		}
	}
	return append(append(lt, eq...), gt...)
}

func main() {
	fmt.Println(pivotArray([]int{9, 12, 5, 10, 14, 3, 10}, 10))
}
