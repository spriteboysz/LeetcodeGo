/**
 * Author: Deean
 * Date: 2022/12/4 12:00
 * FileName: algorithm/P0257. 二叉树的所有路径.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
	"strconv"
)

func binaryTreePaths(root *lib.TreeNode) []string {
	paths := []string{}
	var dfs func(root *lib.TreeNode, path string)
	dfs = func(root *lib.TreeNode, path string) {
		if root == nil {
			return
		}
		pathSB := path
		pathSB += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			paths = append(paths, pathSB)
		} else {
			pathSB += "->"
			dfs(root.Left, pathSB)
			dfs(root.Right, pathSB)
		}
	}
	dfs(root, "")
	return paths
}

func main() {
	fmt.Println(binaryTreePaths(lib.Str2Tree("[1,2,3,null,5]")))
}
