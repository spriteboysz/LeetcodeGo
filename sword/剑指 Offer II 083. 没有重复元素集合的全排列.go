/**
 * Author: Deean
 * Date: 2022/12/5 23:08
 * FileName: sword/剑指 Offer II 083. 没有重复元素集合的全排列.go
 * Description:
 */

package main

import "fmt"

func permute(nums []int) [][]int {
	n, paths, path := len(nums), [][]int{}, []int{}
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
	fmt.Println(permute([]int{1, 2, 3}))
}
