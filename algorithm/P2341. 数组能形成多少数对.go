/**
 * Author: Deean
 * Date: 2022-10-16 17:15
 * FileName: algorithm/P2341. 数组能形成多少数对.go
 * Description:
 */

package main

import "fmt"

func numberOfPairs(nums []int) []int {
	hash := make(map[int]int)
	for _, num := range nums {
		hash[num]++
	}
	pair, single := 0, 0
	for _, v := range hash {
		pair += v / 2
		single += v % 2
	}
	return []int{pair, single}
}

func main() {
	nums := []int{1, 3, 2, 1, 3, 2, 2}
	fmt.Println(numberOfPairs(nums))
}
