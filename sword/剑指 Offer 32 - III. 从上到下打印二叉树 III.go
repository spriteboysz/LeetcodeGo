/**
 * Author: Deean
 * Date: 2022/12/11 17:06
 * FileName: sword/剑指 Offer 32 - III. 从上到下打印二叉树 III.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func levelOrderIII(root *lib.TreeNode) [][]int {
	levels := [][]int{}
	if root == nil {
		return levels
	}
	queue := []*lib.TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[i]
			if len(levels)%2 == 0 {
				level[i] = node.Val
			} else {
				level[size-1-i] = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levels = append(levels, level)
		queue = queue[size:]
	}
	return levels
}

func main() {
	fmt.Println(levelOrderIII(lib.Str2Tree("[3,9,20,null,null,15,7]")))
}
