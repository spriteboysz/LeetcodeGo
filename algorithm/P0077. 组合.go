/**
 * Author: Deean
 * Date: 2022/12/6 23:16
 * FileName: algorithm/P0077. 组合.go
 * Description:
 */

package main

import "fmt"

func combine(n int, k int) [][]int {
	paths, path := [][]int{}, []int{}
	var backtrace func(int)
	backtrace = func(index int) {
		if len(path) == k {
			paths = append(paths, append([]int{}, path...))
			return
		}
		if index > n {
			return
		}
		backtrace(index + 1)
		path = append(path, index)
		backtrace(index + 1)
		path = path[:len(path)-1]
	}

	backtrace(1)
	return paths
}

func main() {
	fmt.Println(combine(4, 2))
}
