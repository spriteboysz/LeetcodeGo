/**
 * Author: Deean
 * Date: 2023-10-16 23:46
 * FileName: LCR/LCR 150. 彩灯装饰记录 II.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func decorateRecord2(root *lib.TreeNode) [][]int {
	type TreeNode = lib.TreeNode
	levels := [][]int{}
	if root == nil {
		return levels
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		level, n := []int{}, len(queue)
		for i := 0; i < n; i++ {
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
	fmt.Println(decorateRecord2(lib.Str2Tree("[8,17,21,18,null,null,6]")))
}
