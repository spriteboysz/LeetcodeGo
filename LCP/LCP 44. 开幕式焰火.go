/**
 * Author: Deean
 * Date: 2022-10-16 17:02
 * FileName: LCP/LCP 44. 开幕式焰火.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func numColor(root *lib.TreeNode) int {
	color := map[int]bool{}
	var dfs func(node *lib.TreeNode)
	dfs = func(root *lib.TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		color[root.Val] = true
		dfs(root.Right)
	}
	dfs(root)
	return len(color)
}

func main() {
	root := lib.Str2Tree("[1,3,2,1,null,2]")
	fmt.Println(numColor(root))
}
