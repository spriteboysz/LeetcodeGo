/**
 * Author: Deean
 * Date: 2022-11-08 23:52
 * FileName: algorithm/P0107. 二叉树的层序遍历 II.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func levelOrderBottom(root *lib.TreeNode) [][]int {
	levels := [][]int{}
	if root == nil {
		return levels
	}
	queue := []*lib.TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		level := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levels = append([][]int{level}, levels...)
	}
	return levels
}

func main() {
	fmt.Println(levelOrderBottom(lib.Str2Tree("[3,9,20,null,null,15,7]")))
}
