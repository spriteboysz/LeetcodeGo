/**
 * Author: Deean
 * Date: 2022-10-16 23:38
 * FileName: algorithm/P0700. 二叉搜索树中的搜索.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func searchBST(root *lib.TreeNode, val int) *lib.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	} else if root.Val < val {
		return searchBST(root.Right, val)
	} else {
		return root
	}
}

func main() {
	root := lib.Str2Tree("[4,2,7,1,3]")
	fmt.Println(lib.Tree2String(searchBST(root, 2)))
	fmt.Println(lib.Tree2String(searchBST(root, 6)))
}
