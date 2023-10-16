/**
 * Author: Deean
 * Date: 2023-10-15 22:38
 * FileName: LCP/LCR 060. 前 K 个高频元素.go
 * Description:
 */

package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	hash := map[int]int{}
	for _, num := range nums {
		hash[num]++
	}
	s := make([]int, len(nums)+1)
	for _, v := range hash {
		s[v]++
	}
	i, t := len(nums), 0
	for t < k {
		t += s[i]
		i -= 1
	}
	ans := make([]int, 0)
	for key, v := range hash {
		if v > i {
			ans = append(ans, key)
		}
	}
	return ans
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
