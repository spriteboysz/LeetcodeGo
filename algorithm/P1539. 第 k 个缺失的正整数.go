/**
 * Author: Deean
 * Date: 2022/11/27 9:25
 * FileName: algorithm/P1539. 第 k 个缺失的正整数.go
 * Description:
 */

package main

import "fmt"

func findKthPositive(arr []int, k int) int {
	for _, num := range arr {
		if num <= k {
			k++
		}
	}
	return k
}

func main() {
	fmt.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
}
