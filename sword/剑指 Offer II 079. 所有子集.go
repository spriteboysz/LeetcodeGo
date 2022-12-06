/**
 * Author: Deean
 * Date: 2022/12/5 23:25
 * FileName: sword/剑指 Offer II 079. 所有子集.go
 * Description:
 */

package main

import "fmt"

func subsets(nums []int) [][]int {
	path, paths := []int{}, [][]int{}
	var backtrace func([]int, int, []int)
	backtrace = func(nums []int, index int, path []int) {
		if index <= len(nums) {
			paths = append(paths, append([]int{}, path...))
		}
		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrace(nums, i+1, path)
			path = path[:len(path)-1]
		}
	}

	backtrace(nums, 0, path)
	return paths
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
