/**
 * Author: Deean
 * Date: 2022-11-09 23:17
 * FileName: algorithm/P0199. 二叉树的右视图.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func rightSideView(root *lib.TreeNode) []int {
	right := []int{}
	if root == nil {
		return right
	}
	queue := []*lib.TreeNode{root}
	for len(queue) > 0 {
		level, size := 0, len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			level = node.Val
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		right = append(right, level)
	}
	return right
}

func main() {
	fmt.Println(rightSideView(lib.Str2Tree("[1,2,3,null,5,null,4]")))
}
