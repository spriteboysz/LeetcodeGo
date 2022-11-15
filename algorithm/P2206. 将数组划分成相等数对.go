/**
 * Author: Deean
 * Date: 2022/11/15 22:39
 * FileName: algorithm/P2206. 将数组划分成相等数对.go
 * Description:
 */

package main

import "fmt"

func divideArray(nums []int) bool {
	hash := map[int]int{}
	for _, num := range nums {
		hash[num] += 1
	}
	for _, v := range hash {
		if v%2 == 1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(divideArray([]int{3, 2, 3, 2, 2, 2}))
}
