/**
 * Author: Deean
 * Date: 2022/12/10 22:35
 * FileName: algorithm/P2150. 找出数组中的所有孤独数字.go
 * Description:
 */

package main

import "fmt"

func findLonely(nums []int) []int {
	hash := map[int]int{}
	for _, num := range nums {
		hash[num]++
	}
	lonely := []int{}
	for k, v := range hash {
		if v == 1 && hash[k-1] == 0 && hash[k+1] == 0 {
			lonely = append(lonely, k)
		}
	}
	return lonely
}

func main() {
	fmt.Println(findLonely([]int{1, 3, 5, 3}))
}
