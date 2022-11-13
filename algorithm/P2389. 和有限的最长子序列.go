/**
 * Author: Deean
 * Date: 2022/11/13 23:07
 * FileName: algorithm/P2389. 和有限的最长子序列.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	res := []int{}
	sort.Ints(nums)
	for _, query := range queries {
		sum := 0
		for j, num := range nums {
			sum += num
			if sum == query {
				res = append(res, j+1)
				break
			}
			if sum > query {
				res = append(res, j)
				break
			}
			if sum < query && j == len(nums)-1 {
				res = append(res, j+1)
			}
		}
	}
	return res
}

func main() {
	fmt.Println(answerQueries([]int{4, 5, 2, 1}, []int{3, 10, 21}))
}
