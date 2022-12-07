/**
 * Author: Deean
 * Date: 2022/12/6 23:26
 * FileName: interview/面试题 08.04. 幂集.go
 * Description:
 */

package main

import "fmt"

func subsets(nums []int) [][]int {
	paths, path, n := [][]int{}, []int{}, len(nums)
	var backtrace func(int, []int)
	backtrace = func(index int, path []int) {
		if index <= n {
			paths = append(paths, append([]int{}, path...))
		}
		for i := index; i < n; i++ {
			path = append(path, nums[i])
			backtrace(i+1, path)
			path = path[:len(path)-1]
		}
	}

	backtrace(0, path)
	return paths
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
