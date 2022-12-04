/**
 * Author: Deean
 * Date: 2022/12/4 11:50
 * FileName: sword/剑指 Offer 68 - II. 二叉树的最近公共祖先.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func lowestCommonAncestor(root, p, q *lib.TreeNode) *lib.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func main() {
	fmt.Println(lowestCommonAncestor(lib.Str2Tree("[3,5,1,6,2,0,8,null,null,7,4]"),
		lib.Str2Tree("[5]"), lib.Str2Tree("[1]")))
}
