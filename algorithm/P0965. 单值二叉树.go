/**
 * Author: Deean
 * Date: 2022-10-19 23:36
 * FileName: algorithm/P0965. 单值二叉树.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func isUnivalTree(root *lib.TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		if root.Val != root.Left.Val || !isUnivalTree(root.Left) {
			return false
		}
	}
	if root.Right != nil {
		if root.Val != root.Right.Val || !isUnivalTree(root.Right) {
			return false
		}
	}
	return true
}

func main() {
	root := lib.Str2Tree("[2,2,2,5,2]")
	fmt.Println(isUnivalTree(root))
}
