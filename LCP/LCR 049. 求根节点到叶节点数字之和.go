/**
 * Author: Deean
 * Date: 2023-10-15 20:39
 * FileName: LCP/LCR 049. 求根节点到叶节点数字之和.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func sumNumbers(root *lib.TreeNode) int {
	type TreeNode = lib.TreeNode
	var dfs func(*TreeNode, int) int
	dfs = func(root *TreeNode, sum int) int {
		if root == nil {
			return 0
		}
		sum = sum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			return sum
		}
		return dfs(root.Left, sum) + dfs(root.Right, sum)
	}
	return dfs(root, 0)
}

func main() {
	fmt.Println(sumNumbers(lib.Str2Tree("[1,2,3]")))
}
