/**
 * Author: Deean
 * Date: 2022/12/4 17:53
 * FileName: algorithm/P0606. 根据二叉树创建字符串.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
	"strconv"
)

func tree2str(root *lib.TreeNode) string {
	if root == nil {
		return ""
	}
	if root.Left == nil && root.Right == nil {
		return strconv.Itoa(root.Val)
	}
	if root.Right == nil {
		return fmt.Sprintf("%d(%s)", root.Val, tree2str(root.Left))
	}
	return fmt.Sprintf("%d(%s)(%s)", root.Val, tree2str(root.Left), tree2str(root.Right))
}

func main() {
	fmt.Println(tree2str(lib.Str2Tree("[1,2,3,4]")))
}
