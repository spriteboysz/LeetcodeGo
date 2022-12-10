/**
 * Author: Deean
 * Date: 2022/12/10 22:41
 * FileName: sword/剑指 Offer 55 - II. 平衡二叉树.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func isBalanced(root *lib.TreeNode) bool {
	if root == nil {
		return true
	}
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
	var high func(node *lib.TreeNode) int
	high = func(root *lib.TreeNode) int {
		if root == nil {
			return 0
		}
		return max(high(root.Left), high(root.Right)) + 1
	}

	return abs(high(root.Left)-high(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func main() {
	fmt.Println(isBalanced(lib.Str2Tree("[3,9,20,null,null,15,7]")))
}
