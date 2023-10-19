/**
 * Author: Deean
 * Date: 2023-10-16 23:18
 * FileName: LCP/LCR 144. 翻转二叉树.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func mirrorTree(root *lib.TreeNode) *lib.TreeNode {
	if root == nil {
		return root
	}
	node := mirrorTree(root.Left)
	root.Left = mirrorTree(root.Right)
	root.Right = node
	return root
}

func main() {
	fmt.Println(lib.Tree2String(mirrorTree(lib.Str2Tree("[5,7,9,8,3,2,4]"))))
}
