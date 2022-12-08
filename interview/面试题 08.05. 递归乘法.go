/**
 * Author: Deean
 * Date: 2022/12/7 23:28
 * FileName: interview/面试题 08.05. 递归乘法.go
 * Description:
 */

package main

import "fmt"

func multiply(A int, B int) int {
	if A < B {
		A, B = B, A
	}
	var dfs func(x int) int
	dfs = func(x int) int {
		if B == 1 {
			return x
		}
		B--
		return x + dfs(x)
	}
	return dfs(A)
}

func main() {
	fmt.Println(multiply(10, 5))
}
