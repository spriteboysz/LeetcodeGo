/**
 * Author: Deean
 * Date: 2022/11/15 23:50
 * FileName: algorithm/P0102. 二叉树的层序遍历.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func levelOrder(root *lib.TreeNode) [][]int {
	levels := [][]int{}
	if root == nil {
		return levels
	}
	queue := []*lib.TreeNode{root}
	for len(queue) > 0 {
		level, size := []int{}, len(queue)
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
		levels = append(levels, level)
	}
	return levels
}

func main() {
	fmt.Println(levelOrder(lib.Str2Tree("[3,9,20,null,null,15,7]")))
}
