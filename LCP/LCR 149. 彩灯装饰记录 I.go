/**
 * Author: Deean
 * Date: 2023-10-16 23:39
 * FileName: algorithm/LCR 149. 彩灯装饰记录 I.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func decorateRecord(root *lib.TreeNode) []int {
	type TreeNode = lib.TreeNode
	levels := []int{}
	if root == nil {
		return levels
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		for i, n := 0, len(queue); i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			levels = append(levels, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return levels
}

func main() {
	fmt.Println(decorateRecord(lib.Str2Tree("[8,17,21,18,null,null,6]")))
}
