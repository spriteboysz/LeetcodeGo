/**
 * Author: Deean
 * Date: 2022-10-14 22:39
 * FileName: lib/TreeNode.go
 * Description:
 */

package lib

import (
	"strconv"
	"strings"
)

// TreeNode is tree's node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NULL 方便添加测试数据
var NULL = -1 << 63

// Str2Tree 利用 []int 生成 *TreeNode
func Str2Tree(s string) *TreeNode {
	ss := strings.Split(s[1:len(s)-1], ",")
	var nums []int
	for _, str := range ss {
		if str == "NULL" || str == "null" {
			nums = append(nums, NULL)
		} else {
			num, _ := strconv.Atoi(str)
			nums = append(nums, num)
		}
	}
	n := len(nums)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: nums[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && nums[i] != NULL {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && nums[i] != NULL {
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// Equal return ture if tn == a
func (tn *TreeNode) Equal(a *TreeNode) bool {
	if tn == nil && a == nil {
		return true
	}

	if tn == nil || a == nil || tn.Val != a.Val {
		return false
	}

	return tn.Left.Equal(a.Left) && tn.Right.Equal(a.Right)
}

// Tree2String 把 *TreeNode 按照行还原成 []int
func Tree2String(tn *TreeNode) string {
	res := make([]string, 0, 1024)

	queue := []*TreeNode{tn}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			nd := queue[i]
			if nd == nil {
				res = append(res, "NULL")
			} else {
				res = append(res, strconv.Itoa(nd.Val))
				queue = append(queue, nd.Left, nd.Right)
			}
		}
		queue = queue[size:]
	}

	i := len(res)
	for i > 0 && res[i-1] == "NULL" {
		i--
	}
	return "[" + strings.Join(res[:i], ",") + "]"
}
