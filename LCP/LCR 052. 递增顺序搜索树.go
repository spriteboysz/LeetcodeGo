/**
 * Author: Deean
 * Date: 2023-10-15 20:53
 * FileName: LCP/LCR 052. 递增顺序搜索树.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func increasingBST(root *lib.TreeNode) *lib.TreeNode {
	type TreeNode = lib.TreeNode
	var nodes []*TreeNode
	var dfs func(node *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		nodes = append(nodes, root)
		dfs(root.Right)
	}

	dfs(root)
	for i, n := 0, len(nodes); i < n-1; i++ {
		nodes[i].Left = nil
		nodes[i].Right = nodes[i+1]
	}
	fmt.Println(len(nodes))
	nodes[len(nodes)-1].Left = nil
	nodes[len(nodes)-1].Right = nil
	return nodes[0]
}

func main() {
	fmt.Println(
		lib.Tree2String(increasingBST(lib.Str2Tree("[5,3,6,2,4,null,8,1,null,null,null,7,9]"))))
}
