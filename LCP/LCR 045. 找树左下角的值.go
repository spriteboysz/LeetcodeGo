/**
 * Author: Deean
 * Date: 2023-10-15 20:24
 * FileName: LCP/LCR 045. 找树左下角的值.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func findBottomLeftValue(root *lib.TreeNode) int {
	type TreeNode = lib.TreeNode
	target := root.Val
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		for i, n := 0, len(queue); i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == 0 {
				target = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return target
}

func main() {
	fmt.Println(findBottomLeftValue(lib.Str2Tree("[1,2,3,4,null,5,6,null,null,7]")))
}
