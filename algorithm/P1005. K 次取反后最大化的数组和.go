/**
 * Author: Deean
 * Date: 2022/11/21 23:17
 * FileName: algorithm/P1005. K 次取反后最大化的数组和.go
 * Description:
 */

package main

import "fmt"

func largestSumAfterKNegations(nums []int, k int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	hash := map[int]int{}
	sum := 0
	for _, num := range nums {
		hash[num]++
		sum += num
	}
	for i := -100; i < 0 && k != 0; i++ {
		if hash[i] > 0 {
			ops := min(hash[i], k)
			sum -= i * ops * 2
			hash[-i] += ops
			k -= ops
		}
	}
	if k > 0 && k%2 == 1 && hash[0] == 0 {
		for i := 1; i <= 100; i++ {
			if hash[i] > 0 {
				sum -= i * 2
				break
			}
		}
	}
	return sum
}

func main() {
	fmt.Println(largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2))
}
