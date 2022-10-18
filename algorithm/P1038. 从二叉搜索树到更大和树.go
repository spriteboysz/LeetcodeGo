/**
 * Author: Deean
 * Date: 2022-10-18 23:04
 * FileName: algorithm/P1038. 从二叉搜索树到更大和树.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func bstToGst(root *lib.TreeNode) *lib.TreeNode {
	sum := 0
	var dfs func(root *lib.TreeNode)
	dfs = func(root *lib.TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		sum += root.Val
		root.Val = sum
		dfs(root.Left)
	}
	dfs(root)
	return root
}

func main() {
	root := lib.Str2Tree("[3,2,4,1]")
	fmt.Println(lib.Tree2String(bstToGst(root)))
}
