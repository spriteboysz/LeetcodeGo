/**
 * Author: Deean
 * Date: 2022-11-08 23:36
 * FileName: sword/剑指 Offer II 056. 二叉搜索树中两个节点之和.go
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
	fmt.Println(findTarget(lib.Str2Tree("[8,6,10,5,7,9,11]"), 22))
}
