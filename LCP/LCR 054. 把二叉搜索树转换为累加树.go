/**
 * Author: Deean
 * Date: 2023-10-15 22:29
 * FileName: LCP/LCR 054. 把二叉搜索树转换为累加树.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func convertBST(root *lib.TreeNode) *lib.TreeNode {
	type TreeNode = lib.TreeNode
	acc := 0
	var dfs func(node *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		acc += root.Val
		root.Val = acc
		dfs(root.Left)
	}

	dfs(root)
	return root
}

func main() {
	fmt.Println(lib.Tree2String(convertBST(lib.Str2Tree("[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]"))))
}
