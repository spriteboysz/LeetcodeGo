/**
 * Author: Deean
 * Date: 2022-10-17 23:55
 * FileName: algorithm/P0094. 二叉树的中序遍历.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func inorderTraversal(root *lib.TreeNode) []int {
	var values []int
	var dfs func(root *lib.TreeNode)
	dfs = func(root *lib.TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		values = append(values, root.Val)
		dfs(root.Right)
	}

	dfs(root)
	return values
}

func main() {
	root := lib.Str2Tree("[1,null,2,3]")
	fmt.Println(inorderTraversal(root))
}
