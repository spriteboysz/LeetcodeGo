/**
 * Author: Deean
 * Date: 2023-10-16 23:51
 * FileName: LCR/LCR 151. 彩灯装饰记录 III.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func decorateRecord3(root *lib.TreeNode) [][]int {
	type TreeNode = lib.TreeNode
	levels := [][]int{}
	if root == nil {
		return levels
	}
	queue, depth := []*TreeNode{}, 0
	queue = append(queue, root)
	for len(queue) > 0 {
		level, n := []int{}, len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if depth%2 == 0 {
				level = append(level, node.Val)
			} else {
				level = append([]int{node.Val}, level...)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
		levels = append(levels, level)
	}
	return levels
}

func main() {
	fmt.Println(decorateRecord3(lib.Str2Tree("[8,17,21,18,null,null,6]")))
}
