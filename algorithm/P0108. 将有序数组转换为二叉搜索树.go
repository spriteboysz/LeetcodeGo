/**
 * Author: Deean
 * Date: 2022-10-17 22:56
 * FileName: algorithm/P0108. 将有序数组转换为二叉搜索树.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func sortedArrayToBST(nums []int) *lib.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) >> 1
	root := &lib.TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	fmt.Println(lib.Tree2String(sortedArrayToBST(nums)))
}
