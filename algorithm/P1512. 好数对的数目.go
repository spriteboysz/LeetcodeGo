/**
 * Author: Deean
 * Date: 2022-10-13 23:56
 * FileName: algorithm/P1512. 好数对的数目.go
 * Description:
 */

package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	cnt := 0
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num == nums[j] {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	nums := []int{1, 1, 1, 1}
	fmt.Println(numIdenticalPairs(nums))
}
