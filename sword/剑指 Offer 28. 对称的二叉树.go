/**
 * Author: Deean
 * Date: 2022/11/20 17:30
 * FileName: sword/剑指 Offer 28. 对称的二叉树.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func isSymmetric(root *lib.TreeNode) bool {
	var check func(p, q *lib.TreeNode) bool
	check = func(p, q *lib.TreeNode) bool {
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
	fmt.Println(isSymmetric(lib.Str2Tree("[1,2,2,3,4,4,3]")))
}
