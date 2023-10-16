/**
 * Author: Deean
 * Date: 2023-10-15 20:30
 * FileName: LCP/LCR 046. 二叉树的右视图.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func rightSideView(root *lib.TreeNode) []int {
	type TreeNode = lib.TreeNode
	levels := []int{}
	if root == nil {
		return levels
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		level := queue[0].Val
		for i, n := 0, len(queue); i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			level = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levels = append(levels, level)
	}
	return levels
}

func main() {
	fmt.Println(rightSideView(lib.Str2Tree("[1,2,3,null,5,null,4]")))
}
