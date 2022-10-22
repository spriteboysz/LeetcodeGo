/**
 * Author: Deean
 * Date: 2022-10-22 16:22
 * FileName: algorithm/P1302. 层数最深叶子节点的和.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func deepestLeavesSum(root *lib.TreeNode) int {
	sum := 0
	queue := []*lib.TreeNode{root}
	for len(queue) > 0 {
		sum = 0
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return sum
}

func main() {
	root := lib.Str2Tree("[6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]")
	fmt.Println(deepestLeavesSum(root))
}
