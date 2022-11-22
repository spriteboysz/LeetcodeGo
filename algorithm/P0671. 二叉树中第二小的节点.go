/**
 * Author: Deean
 * Date: 2022/11/22 23:10
 * FileName: algorithm/P0671. 二叉树中第二小的节点.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
	"sort"
)

func findSecondMinimumValue(root *lib.TreeNode) int {
	hash := map[int]bool{}
	var dfs func(root *lib.TreeNode)
	dfs = func(root *lib.TreeNode) {
		if root == nil {
			return
		}
		hash[root.Val] = true
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	values := []int{}
	for k := range hash {
		values = append(values, k)
	}
	if len(values) <= 1 {
		return -1
	}
	sort.Ints(values)
	return values[1]
}

func main() {
	fmt.Println(findSecondMinimumValue(lib.Str2Tree("[2,2,5,null,null,5,7]")))
}
