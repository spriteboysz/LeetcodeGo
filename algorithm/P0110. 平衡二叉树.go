/**
 * Author: Deean
 * Date: 2022/12/10 22:52
 * FileName: algorithm/P0110. 平衡二叉树.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func isBalanced(root *lib.TreeNode) bool {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	abs := func(a int) int {
		if a >= 0 {
			return a
		}
		return -a
	}
	var high func(*lib.TreeNode) int
	high = func(root *lib.TreeNode) int {
		if root == nil {
			return 0
		}
		return max(high(root.Left), high(root.Right)) + 1
	}

	if root == nil {
		return true
	}
	return abs(high(root.Left)-high(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func main() {
	fmt.Println(isBalanced(lib.Str2Tree("[3,9,20,null,null,15,7]")))
}
