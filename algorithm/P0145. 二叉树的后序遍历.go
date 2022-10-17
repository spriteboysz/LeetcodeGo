/**
 * Author: Deean
 * Date: 2022-10-17 23:52
 * FileName: algorithm/P0145. 二叉树的后序遍历.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func postorderTraversal(root *lib.TreeNode) []int {
	var values []int
	var dfs func(root *lib.TreeNode)
	dfs = func(root *lib.TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		values = append(values, root.Val)
	}

	dfs(root)
	return values
}

func main() {
	root := lib.Str2Tree("[1,null,2,3]")
	fmt.Println(postorderTraversal(root))
}
