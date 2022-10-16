/**
 * Author: Deean
 * Date: 2022-10-16 20:44
 * FileName: interview/面试题 04.02. 最小高度树.go
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
