/**
 * Author: Deean
 * Date: 2022-10-21 22:41
 * FileName: algorithm/P0268. 丢失的数字.go
 * Description:
 */

package main

import "fmt"

func missingNumber(nums []int) int {
	total, n := 0, len(nums)
	for _, num := range nums {
		total += num
	}
	return n*(n+1)/2 - total
}

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(nums))
}
