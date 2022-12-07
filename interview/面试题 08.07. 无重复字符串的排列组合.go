/**
 * Author: Deean
 * Date: 2022/12/6 23:37
 * FileName: sword/面试题 08.07. 无重复字符串的排列组合.go
 * Description:
 */

package main

import "fmt"

func permutation(S string) []string {
	paths, path, n := []string{}, []byte{}, len(S)
	visited := make([]bool, n)
	var backtrace func(int)
	backtrace = func(index int) {
		if index == n {
			paths = append(paths, string(path))
		}
		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			path = append(path, S[i])
			visited[i] = true
			backtrace(index + 1)
			visited[i] = false
			path = path[:len(path)-1]
		}
	}

	backtrace(0)
	return paths
}

func main() {
	fmt.Println(permutation("qwe"))
}
