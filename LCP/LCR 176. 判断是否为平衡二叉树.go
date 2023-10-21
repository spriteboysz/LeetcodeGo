/**
 * Author: Deean
 * Date: 2023-10-19 23:10
 * FileName: LCR/LCR 176. 判断是否为平衡二叉树.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func isBalanced(root *lib.TreeNode) bool {
	type TreeNode = lib.TreeNode
	Abs := func(a int) int {
		if a >= 0 {
			return a
		}
		return -a
	}

	Max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var depth func(*TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return Max(depth(root.Left), depth(root.Right)) + 1
	}

	if root == nil {
		return true
	}
	return Abs(depth(root.Left)-depth(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func main() {
	fmt.Println(isBalanced(lib.Str2Tree("[3,9,20,null,null,15,7]")))
}
