/**
 * Author: Deean
 * Date: 2022-10-15 23:32
 * FileName: algorithm/P0938. 二叉搜索树的范围和.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func rangeSumBST(root *lib.TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

func main() {
	root := lib.Str2Tree("[10,5,15,3,7,13,18,1,null,6]")
	fmt.Println(rangeSumBST(root, 6, 10))
}
