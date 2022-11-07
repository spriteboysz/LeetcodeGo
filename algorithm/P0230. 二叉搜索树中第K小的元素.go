/**
 * Author: Deean
 * Date: 2022-11-07 23:19
 * FileName: algorithm/P0230. 二叉搜索树中第K小的元素.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func kthSmallest(root *lib.TreeNode, k int) int {
	var values []int
	var dfs func(root *lib.TreeNode)
	dfs = func(root *lib.TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		values = append(values, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return values[k-1]
}

func main() {
	root := lib.Str2Tree("[5,3,6,2,4,null,null,1]")
	fmt.Println(kthSmallest(root, 3))
}
