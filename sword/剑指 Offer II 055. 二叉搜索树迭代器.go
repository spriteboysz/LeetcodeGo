/**
 * Author: Deean
 * Date: 2022-10-22 16:29
 * FileName: sword/剑指 Offer II 055. 二叉搜索树迭代器.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

type BSTIterator struct {
	values []int
}

func Constructor(root *lib.TreeNode) BSTIterator {
	it := BSTIterator{}
	var dfs func(root *lib.TreeNode)
	dfs = func(root *lib.TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		it.values = append(it.values, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return it
}

func (it *BSTIterator) Next() int {
	val := it.values[0]
	it.values = it.values[1:]
	return val
}

func (it *BSTIterator) HasNext() bool {
	return len(it.values) > 0
}

func main() {
	root := lib.Str2Tree("[7, 3, 15, null, null, 9, 20]")
	bSTIterator := Constructor(root)
	fmt.Println(
		bSTIterator.Next(),    // 返回 3
		bSTIterator.Next(),    // 返回 7
		bSTIterator.HasNext(), // 返回 True
		bSTIterator.Next(),    // 返回 9
		bSTIterator.HasNext(), // 返回 True
		bSTIterator.Next(),    // 返回 15
		bSTIterator.HasNext(), // 返回 True
		bSTIterator.Next(),    // 返回 20
		bSTIterator.HasNext(), // 返回 False
	)
}
