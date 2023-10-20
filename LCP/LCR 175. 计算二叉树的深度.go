/**
 * Author: Deean
 * Date: 2023-10-19 23:07
 * FileName: LCR/LCR 175. 计算二叉树的深度.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func calculateDepth(root *lib.TreeNode) int {
	Max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	if root == nil {
		return 0
	}
	return Max(calculateDepth(root.Left), calculateDepth(root.Right)) + 1
}

func main() {
	fmt.Println(calculateDepth(lib.Str2Tree("[1, 2, 2, 3, null, null, 5, 4, null, null, 4]")))
}
