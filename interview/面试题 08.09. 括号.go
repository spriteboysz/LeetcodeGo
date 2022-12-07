/**
 * Author: Deean
 * Date: 2022/12/6 23:30
 * FileName: interview/面试题 08.09. 括号.go
 * Description:
 */

package main

import "fmt"

func generateParenthesis(n int) []string {
	paths, path := []string{}, []byte{}
	left, right := 0, 0
	var backtrace func(int)
	backtrace = func(cnt int) {
		if cnt == 2*n {
			paths = append(paths, string(path))
			return
		}
		if left < n {
			left++
			path = append(path, '(')
			backtrace(cnt + 1)
			left--
			path = path[:len(path)-1]
		}
		if left > right {
			right++
			path = append(path, ')')
			backtrace(cnt + 1)
			right--
			path = path[:len(path)-1]
		}
	}

	backtrace(0)
	return paths
}

func main() {
	fmt.Println(generateParenthesis(3))
}
