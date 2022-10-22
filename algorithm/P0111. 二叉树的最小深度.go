/**
 * Author: Deean
 * Date: 2022-10-22 15:44
 * FileName: algorithm/P0111. 二叉树的最小深度.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
	"math"
)

func minDepth(root *lib.TreeNode) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minimum := math.MaxInt32
	if root.Left != nil {
		minimum = min(minimum, minDepth(root.Left))
	}
	if root.Right != nil {
		minimum = min(minimum, minDepth(root.Right))
	}
	return minimum + 1
}

func main() {
	root := lib.Str2Tree("[2,null,3,null,4,null,5,null,6]")
	fmt.Println(minDepth(root))
}
