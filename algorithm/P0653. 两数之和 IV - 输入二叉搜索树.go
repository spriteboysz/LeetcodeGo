/**
 * Author: Deean
 * Date: 2022-11-08 23:46
 * FileName: algorithm/P0653. 两数之和 IV - 输入二叉搜索树.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func findTarget(root *lib.TreeNode, k int) bool {
	values := map[int]bool{}
	var dfs func(root *lib.TreeNode)
	dfs = func(root *lib.TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		values[root.Val] = true
		dfs(root.Right)
	}
	dfs(root)
	for val := range values {
		if values[k-val] && k-val != val {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(findTarget(lib.Str2Tree("[5,3,6,2,4,null,7]"), 9))
}
