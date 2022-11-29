/**
 * Author: Deean
 * Date: 2022/11/29 23:54
 * FileName: algorithm/P0513. 找树左下角的值.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func findBottomLeftValue(root *lib.TreeNode) int {
	target := 0
	queue := []*lib.TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		target = node.Val
	}
	return target
}

func main() {
	fmt.Println(findBottomLeftValue(lib.Str2Tree("[1,2,3,4,null,5,6,null,null,7]")))
}
