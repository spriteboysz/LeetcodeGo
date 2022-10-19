/**
 * Author: Deean
 * Date: 2022-10-20 00:07
 * FileName: algorithm/P0144. 二叉树的前序遍历.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func preorderTraversal(root *lib.TreeNode) []int {
	var values []int
	var dfs func(root *lib.TreeNode)
	dfs = func(root *lib.TreeNode) {
		if root == nil {
			return
		}
		values = append(values, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return values
}

func main() {
	root := lib.Str2Tree("[1,null,2,3]")
	fmt.Println(preorderTraversal(root))
}
