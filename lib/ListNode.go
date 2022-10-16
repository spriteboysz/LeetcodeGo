/**
 * Author: Deean
 * Date: 2022-10-14 22:45
 * FileName: lib/ListNode.go
 * Description:
 */

package lib

import (
	"strconv"
	"strings"
)

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// List2String convert List to string
func List2String(head *ListNode) string {
	var res []string
	for head != nil {
		res = append(res, strconv.Itoa(head.Val))
		head = head.Next
	}

	return "[" + strings.Join(res, ",") + "]"
}

// Str2List convert []int to List
func Str2List(s string) *ListNode {
	strs := strings.Split(s[1:len(s)-1], ",")
	var nums []int
	for _, str := range strs {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}
	if len(nums) == 0 {
		return nil
	}

	dummy := &ListNode{}
	cur := dummy
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

// GetNodeWith returns the first node with val
func (l *ListNode) GetNodeWith(val int) *ListNode {
	res := l
	for res != nil {
		if res.Val == val {
			break
		}
		res = res.Next
	}
	return res
}

// Str2ListWithCycle returns a list whose tail point to pos-indexed node
// head's index is 0
// if pos = -1, no cycle
func Str2ListWithCycle(s string, pos int) *ListNode {
	head := Str2List(s)
	if pos == -1 {
		return head
	}
	c := head
	for pos > 0 {
		c = c.Next
		pos--
	}
	tail := c
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = c
	return head
}
