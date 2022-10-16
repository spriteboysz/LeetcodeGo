/**
 * Author: Deean
 * Date: 2022-10-16 16:15
 * FileName: algorithm/P1480. 一维数组的动态和.go
 * Description:
 */

package main

import "fmt"

func runningSum(nums []int) []int {
	var running []int
	cur := 0
	for _, num := range nums {
		cur += num
		running = append(running, cur)
	}
	return running
}

func main() {
	nums := []int{3, 1, 2, 10, 1}
	fmt.Println(runningSum(nums))
}
