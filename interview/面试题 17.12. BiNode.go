/**
 * Author: Deean
 * Date: 2022/11/18 23:31
 * FileName: interview/面试题 17.12. BiNode.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func convertBiNode(root *lib.TreeNode) *lib.TreeNode {
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
	fmt.Println(lib.Tree2String(convertBiNode(lib.Str2Tree("[4,2,5,1,3,null,6,0]"))))
}
