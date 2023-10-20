/**
 * Author: Deean
 * Date: 2023-10-19 23:03
 * FileName: LCR/LCR 174. 寻找二叉搜索树中的目标节点.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func findTargetNode(root *lib.TreeNode, cnt int) int {
	type TreeNode = lib.TreeNode
	nodes := []int{}

	var dfs func(node *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		nodes = append(nodes, root.Val)
		dfs(root.Right)
	}

	dfs(root)
	return nodes[len(nodes)-cnt]
}

func main() {
	fmt.Println(findTargetNode(lib.Str2Tree("[7, 3, 9, 1, 5]"), 2))
}
