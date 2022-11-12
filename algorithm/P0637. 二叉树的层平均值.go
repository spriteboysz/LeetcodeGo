/**
 * Author: Deean
 * Date: 2022/11/12 16:18
 * FileName: algorithm/P0637. 二叉树的层平均值.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func averageOfLevels(root *lib.TreeNode) []float64 {
	average := []float64{}
	if root == nil {
		return average
	}
	queue := []*lib.TreeNode{root}
	for len(queue) > 0 {
		size, level, n := len(queue), 0, 0
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level += node.Val
			n += 1
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		average = append(average, float64(level*1.0)/float64(n*1.0))
	}
	return average
}

func main() {
	fmt.Println(averageOfLevels(lib.Str2Tree("[3,9,20,null,null,15,7]")))
}
