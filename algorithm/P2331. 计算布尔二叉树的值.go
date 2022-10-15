/**
 * Author: Deean
 * Date: 2022-10-15 22:30
 * FileName: algorithm/P2331. 计算布尔二叉树的值.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func evaluateTree(root *lib.TreeNode) bool {
	if root.Left == nil {
		return root.Val == 1
	}
	if root.Val == 2 {
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	} else {
		return evaluateTree(root.Left) && evaluateTree(root.Right)
	}

}

func main() {
	root := lib.Str2Tree("[2,1,3,null,null,0,1]")
	fmt.Println(evaluateTree(root))
}
