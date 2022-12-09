/**
 * Author: Deean
 * Date: 2022/12/8 23:13
 * FileName: sword/剑指 Offer II 082. 含有重复元素集合的组合.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	paths, path, n := [][]int{}, []int{}, len(candidates)
	var backtrace func(int, int)
	backtrace = func(index, target int) {
		if target == 0 {
			paths = append(paths, append([]int{}, path...))
			return
		}
		if index == n || candidates[index] > target {
			return
		}
		path = append(path, candidates[index])
		backtrace(index+1, target-candidates[index])
		path = path[:len(path)-1]
		for index < n-1 && candidates[index] == candidates[index+1] {
			index++
		}
		backtrace(index+1, target)
	}

	backtrace(0, target)
	return paths
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
