/**
 * Author: Deean
 * Date: 2022/11/22 23:46
 * FileName: sword/剑指 Offer 53 - II. 0～n-1中缺失的数字.go
 * Description:
 */

package main

import "fmt"

func missingNumber(nums []int) int {
	sum, n := 0, len(nums)
	for _, num := range nums {
		sum += num
	}
	return (1+n)*n/2 - sum
}

func main() {
	fmt.Println(missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 9}))
}
