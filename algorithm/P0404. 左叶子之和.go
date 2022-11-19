/**
 * Author: Deean
 * Date: 2022/11/19 15:54
 * FileName: algorithm/P0404. 左叶子之和.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func sumOfLeftLeaves(root *lib.TreeNode) int {
	isLeaf := func(node *lib.TreeNode) bool {
		return node.Left == nil && node.Right == nil
	}

	var dfs func(root *lib.TreeNode) int
	dfs = func(node *lib.TreeNode) (sum int) {
		if node.Left != nil {
			if isLeaf(node.Left) {
				sum += node.Left.Val
			} else {
				sum += dfs(node.Left)
			}
		}
		if node.Right != nil && !isLeaf(node.Right) {
			sum += dfs(node.Right)
		}
		return
	}
	if root == nil {
		return 0
	}
	return dfs(root)
}

func main() {
	fmt.Println(sumOfLeftLeaves(lib.Str2Tree("[3,9,20,null,null,15,7]")))
}
