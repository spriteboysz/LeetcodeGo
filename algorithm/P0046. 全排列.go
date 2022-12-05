/**
 * Author: Deean
 * Date: 2022/12/5 23:18
 * FileName: algorithm/P0046. 全排列.go
 * Description:
 */

package main

import "fmt"

func permute(nums []int) [][]int {
	path, paths, n := []int{}, [][]int{}, len(nums)
	visited := make([]bool, n)
	var backtrace func(cnt int)
	backtrace = func(cnt int) {
		if cnt == n {
			paths = append(paths, append([]int{}, path...))
			return
		}
		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			path = append(path, nums[i])
			visited[i] = true
			backtrace(cnt + 1)
			path = path[:len(path)-1]
			visited[i] = false
		}
	}

	backtrace(0)
	return paths
}

func main() {
	fmt.Println(permute([]int{2, 3, 4}))
}
