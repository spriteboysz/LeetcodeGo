/**
 * Author: Deean
 * Date: 2023-10-15 22:33
 * FileName: LCP/LCR 056. 两数之和 IV - 输入二叉搜索树.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func findTarget(root *lib.TreeNode, k int) bool {
	type TreeNode = lib.TreeNode
	var values []int
	var dfs func(node *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		values = append(values, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	left, right := 0, len(values)-1
	for left < right {
		if values[left]+values[right] < k {
			left++
		} else if values[left]+values[right] > k {
			right--
		} else {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(findTarget(lib.Str2Tree("[8,6,10,5,7,9,11]"), 15))
}
