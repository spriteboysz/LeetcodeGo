/**
 * Author: Deean
 * Date: 2022-10-16 18:04
 * FileName: sword/剑指 Offer 55 - I. 二叉树的深度.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func maxDepth(root *lib.TreeNode) int {
	if root == nil {
		return 0
	}
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func main() {
	root := lib.Str2Tree("[3,9,20,null,null,15,7]")
	fmt.Println(maxDepth(root))
}
