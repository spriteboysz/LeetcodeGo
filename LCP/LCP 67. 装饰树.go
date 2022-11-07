/**
 * Author: Deean
 * Date: 2022-11-06 22:45
 * FileName: LCP/LCP 67. 装饰树.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func expandBinaryTree(root *lib.TreeNode) *lib.TreeNode {
	if root.Left != nil {
		root.Left = &lib.TreeNode{Val: -1, Left: expandBinaryTree(root.Left)}
	}
	if root.Right != nil {
		root.Right = &lib.TreeNode{Val: -1, Right: expandBinaryTree(root.Right)}
	}
	return root
}

func main() {
	root := lib.Str2Tree("[7,5,6]")
	fmt.Println(lib.Tree2String(expandBinaryTree(root)))
}
