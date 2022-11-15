/**
 * Author: Deean
 * Date: 2022/11/15 23:44
 * FileName: sword/剑指 Offer 32 - II. 从上到下打印二叉树 II.go
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
