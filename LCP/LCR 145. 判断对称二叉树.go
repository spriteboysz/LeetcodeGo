/**
 * Author: Deean
 * Date: 2023-10-16 23:21
 * FileName: LCP/LCR 145. 判断对称二叉树.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func checkSymmetricTree(root *lib.TreeNode) bool {
	type TreeNode = lib.TreeNode
	var check func(*TreeNode, *TreeNode) bool
	check = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
	}

	return check(root, root)
}

func main() {
	fmt.Println(checkSymmetricTree(lib.Str2Tree("[6,7,7,8,9,9,8]")))
}
