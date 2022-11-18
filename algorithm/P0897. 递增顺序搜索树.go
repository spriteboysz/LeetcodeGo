/**
 * Author: Deean
 * Date: 2022/11/18 22:38
 * FileName: algorithm/P0897. 递增顺序搜索树.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func increasingBST(root *lib.TreeNode) *lib.TreeNode {
	nodes := []int{}
	var dfs func(root *lib.TreeNode)
	dfs = func(root *lib.TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		nodes = append(nodes, root.Val)
		dfs(root.Right)
	}
	dfs(root)

	dummyNode := &lib.TreeNode{}
	node := dummyNode
	for _, v := range nodes {
		node.Right = &lib.TreeNode{Val: v}
		node = node.Right
	}
	return dummyNode.Right
}

func main() {
	fmt.Println(lib.Tree2String(increasingBST(lib.Str2Tree("[5,1,7]"))))
}
