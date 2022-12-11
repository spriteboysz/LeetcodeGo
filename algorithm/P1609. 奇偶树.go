/**
 * Author: Deean
 * Date: 2022/12/11 17:20
 * FileName: algorithm/P1609. 奇偶树.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func isEvenOddTree(root *lib.TreeNode) bool {
	queue, level := []*lib.TreeNode{root}, 0
	for len(queue) > 0 {
		size, cur := len(queue), 0
		if level%2 == 1 {
			cur = 1000001
		}
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if level%2 == node.Val%2 || (level%2 == 0 && cur >= node.Val) || (level%2 != 0 && cur <= node.Val) {
				return false
			}
			cur = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		level++
	}
	return true
}

func main() {
	fmt.Println(isEvenOddTree(lib.Str2Tree("[1,10,4,3,null,7,9,12,8,6,null,null,2]")))
}
