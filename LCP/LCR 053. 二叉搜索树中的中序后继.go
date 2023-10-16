/**
 * Author: Deean
 * Date: 2023-10-15 22:20
 * FileName: LCP/LCR 053. 二叉搜索树中的中序后继.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func inorderSuccessor(root *lib.TreeNode, p *lib.TreeNode) *lib.TreeNode {
	type TreeNode = lib.TreeNode
	var nodes []*TreeNode
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		nodes = append(nodes, root)
		dfs(root.Right)
	}

	dfs(root)
	for i, node := range nodes[:len(nodes)-1] {
		if node.Val == p.Val {
			return nodes[i+1]
		}
	}
	return nil
}

func main() {
	root := lib.Str2Tree("[5,3,6,2,4,null,null,1]")
	p := lib.Str2Tree("[6]")
	fmt.Println(lib.Tree2String(inorderSuccessor(root, p)))
}
