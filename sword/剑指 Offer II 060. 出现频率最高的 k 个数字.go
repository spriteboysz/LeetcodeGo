/**
 * Author: Deean
 * Date: 2022/12/7 23:12
 * FileName: sword/剑指 Offer II 060. 出现频率最高的 k 个数字.go
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
