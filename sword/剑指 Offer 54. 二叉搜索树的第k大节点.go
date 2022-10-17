/**
 * Author: Deean
 * Date: 2022-10-17 23:37
 * FileName: sword/剑指 Offer 54. 二叉搜索树的第k大节点.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func kthLargest(root *lib.TreeNode, k int) int {
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
	return values[len(values)-k]
}

func main() {
	root := lib.Str2Tree("[5,3,6,2,4,null,null,1]")
	fmt.Println(kthLargest(root, 3))
}
