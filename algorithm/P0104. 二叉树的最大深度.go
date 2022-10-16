/**
 * Author: Deean
 * Date: 2022-10-16 18:07
 * FileName: algorithm/P0104. 二叉树的最大深度.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func maxDepth2(root *lib.TreeNode) int {
	if root == nil {
		return 0
	}
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	return max(maxDepth2(root.Left), maxDepth2(root.Right)) + 1
}

func main() {
	root := lib.Str2Tree("[3,9,20,null,null,15,7]")
	fmt.Println(maxDepth2(root))
}
