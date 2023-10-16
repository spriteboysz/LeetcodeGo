/**
 * Author: Deean
 * Date: 2023-10-15 20:35
 * FileName: LCP/LCR 047. 二叉树剪枝.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func pruneTree(root *lib.TreeNode) *lib.TreeNode {
	// type TreeNode = lib.TreeNode
	if root == nil {
		return root
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}

func main() {
	fmt.Println(lib.Tree2String(pruneTree(lib.Str2Tree("[1,null,0,0,1]"))))
}
