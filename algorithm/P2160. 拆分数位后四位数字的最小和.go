/**
 * Author: Deean
 * Date: 2022-10-13 23:06
 * FileName: algorithm/P2160. 拆分数位后四位数字的最小和.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func minimumSum(num int) int {
	var nums []int
	for num > 0 {
		nums = append(nums, num%10)
		num /= 10
	}
	sort.Ints(nums)
	return 10*(nums[0]+nums[1]) + nums[2] + nums[3]
}

func main() {
	fmt.Println(minimumSum(4009))
}
