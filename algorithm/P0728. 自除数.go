/**
 * Author: Deean
 * Date: 2022-10-16 22:55
 * FileName: algorithm/P0728. 自除数.go
 * Description:
 */

package main

import "fmt"

func selfDividingNumbers(left int, right int) []int {
	isSelfDividing := func(num int) bool {
		number := num
		for num > 0 {
			if num%10 == 0 || number%(num%10) != 0 {
				return false
			}
			num /= 10
		}
		return true
	}
	var nums []int
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			nums = append(nums, i)
		}
	}
	return nums
}

func main() {
	fmt.Println(selfDividingNumbers(47, 85))
}
