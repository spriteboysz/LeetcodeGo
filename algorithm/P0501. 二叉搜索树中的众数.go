/**
 * Author: Deean
 * Date: 2022/11/24 22:51
 * FileName: algorithm/P0501. 二叉搜索树中的众数.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func findMode(root *lib.TreeNode) []int {
	values := map[int]int{}
	var dfs func(root *lib.TreeNode)
	dfs = func(root *lib.TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		values[root.Val]++
		dfs(root.Right)
	}
	dfs(root)
	maximum := 0
	for _, v := range values {
		if v > maximum {
			maximum = v
		}
	}
	mode := []int{}
	for k, v := range values {
		if v == maximum {
			mode = append(mode, k)
		}
	}
	return mode
}

func main() {
	fmt.Println(findMode(lib.Str2Tree("[1,null,2,2]")))
}
