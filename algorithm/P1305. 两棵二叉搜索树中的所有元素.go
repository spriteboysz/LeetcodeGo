/**
 * Author: Deean
 * Date: 2022-10-16 22:46
 * FileName: algorithm/P1305. 两棵二叉搜索树中的所有元素.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
	"sort"
)

func getAllElements(root1 *lib.TreeNode, root2 *lib.TreeNode) []int {
	var values []int
	var dfs = func(root *lib.TreeNode) {}
	dfs = func(root *lib.TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		values = append(values, root.Val)
		dfs(root.Right)
	}

	dfs(root1)
	dfs(root2)
	sort.Ints(values)
	return values
}

func main() {
	root1 := lib.Str2Tree("[2,1,4]")
	root2 := lib.Str2Tree("[1,0,3]")
	fmt.Println(getAllElements(root1, root2))
}
