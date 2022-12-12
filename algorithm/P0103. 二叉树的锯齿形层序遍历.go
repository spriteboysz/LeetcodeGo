/**
 * Author: Deean
 * Date: 2022/12/11 22:37
 * FileName: algorithm/P0103. 二叉树的锯齿形层序遍历.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func zigzagLevelOrder(root *lib.TreeNode) [][]int {
	levels := [][]int{}
	if root == nil {
		return levels
	}
	n := 0
	queue := []*lib.TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0)
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if n&1 == 0 {
				level = append(level, node.Val)
			} else {
				level = append(append([]int{}, node.Val), level...)
			}
			size--
		}
		n++
		levels = append(levels, level)
	}
	return levels
}

func main() {
	fmt.Println(zigzagLevelOrder(lib.Str2Tree("[3,9,20,null,null,15,7]")))
}
