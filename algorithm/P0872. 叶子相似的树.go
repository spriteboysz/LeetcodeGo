/**
 * Author: Deean
 * Date: 2022/11/12 23:12
 * FileName: algorithm/P0872. 叶子相似的树.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func leafSimilar(root1 *lib.TreeNode, root2 *lib.TreeNode) bool {
	values := []int{}
	var dfs func(node *lib.TreeNode)
	dfs = func(node *lib.TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			values = append(values, node.Val)
			return
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root1)
	values1 := append([]int(nil), values...)
	values = []int{}
	dfs(root2)
	values2 := append([]int(nil), values...)
	if len(values2) != len(values1) {
		return false
	}
	for i, v := range values1 {
		if v != values2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(leafSimilar(lib.Str2Tree("[1,2,3]"), lib.Str2Tree("[3,2,1]")))
}
