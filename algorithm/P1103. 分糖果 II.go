/**
 * Author: Deean
 * Date: 2022/11/13 16:41
 * FileName: algorithm/P1103. 分糖果 II.go
 * Description:
 */

package main

import "fmt"

func distributeCandies2(candies int, num_people int) []int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	nums := make([]int, num_people)
	for i := 0; candies > 0; i++ {
		cur := min(i+1, candies)
		nums[i%num_people] += cur
		candies -= cur
	}
	return nums
}

func main() {
	fmt.Println(distributeCandies2(7, 4))
	fmt.Println(distributeCandies2(10, 3))
}
