/**
 * Author: Deean
 * Date: 2022-10-18 23:35
 * FileName: algorithm/P2357. 使数组中所有元素都等于零.go
 * Description:
 */

package main

import "fmt"

func minimumOperations(nums []int) int {
	hash := map[int]int{}
	for _, num := range nums {
		if num > 0 {
			hash[num]++
		}
	}
	return len(hash)
}

func main() {
	nums := []int{1, 5, 0, 3, 5}
	fmt.Println(minimumOperations(nums))
}
