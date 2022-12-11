/**
 * Author: Deean
 * Date: 2022/12/10 23:29
 * FileName: interview/面试题 04.04. 检查平衡性.go
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
	var high func(node *lib.TreeNode) int
	high = func(root *lib.TreeNode) int {
		if root == nil {
			return 0
		}
		return max(high(root.Left), high(root.Right)) + 1
	}

	if root == nil {
		return true
	}
	return abs(abs(high(root.Left)-high(root.Right))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func main() {
	fmt.Println(isBalanced(lib.Str2Tree("[1,2,2,3,3,null,null,4,4]")))
}
