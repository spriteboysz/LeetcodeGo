/**
 * Author: Deean
 * Date: 2022-10-14 21:40
 * FileName: algorithm/P2367. 算术三元组的数目.go
 * Description:
 */

package main

import "fmt"

func arithmeticTriplets(nums []int, diff int) int {
	cnt := 0
	for i, a := range nums {
		for j, b := range nums[i+1:] {
			if b-a != diff {
				continue
			}
			for _, c := range nums[j+1:] {
				if c-b == diff {
					cnt++
				}
			}
		}
	}
	return cnt
}

func main() {
	nums := []int{4, 5, 6, 7, 8, 9}
	fmt.Println(arithmeticTriplets(nums, 2))
}
