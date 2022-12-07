/**
 * Author: Deean
 * Date: 2022/12/7 22:26
 * FileName: sword/剑指 Offer II 081. 允许重复选择元素的组合.go
 * Description:
 */

package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	paths, path := [][]int{}, []int{}
	var backtrace func(int, int)
	backtrace = func(target, index int) {
		if index == len(candidates) {
			return
		}
		if target == 0 {
			paths = append(paths, append([]int{}, path...))
			return
		}
		backtrace(target, index+1)
		if target-candidates[index] >= 0 {
			path = append(path, candidates[index])
			backtrace(target-candidates[index], index)
			path = path[:len(path)-1]
		}
	}

	backtrace(target, 0)
	return paths
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
