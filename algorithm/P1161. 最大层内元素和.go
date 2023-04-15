/**
 * Author: Deean
 * Date: 2022/12/16 22:45
 * FileName: algorithm/P1161. 最大层内元素和.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
	"math"
)

func maxLevelSum(root *lib.TreeNode) int {
	levels := []int{}
	queue := []*lib.TreeNode{root}
	for len(queue) > 0 {
		level, size := 0, len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levels = append(levels, level)
	}
	// fmt.Println(levels)
	maximum := math.MinInt32
	for _, num := range levels {
		if maximum < num {
			maximum = num
		}
	}
	for i, num := range levels {
		if num == maximum {
			return i + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(maxLevelSum(lib.Str2Tree("[1,7,0,7,-8,null,null]")))
}
