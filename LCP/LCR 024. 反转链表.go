/**
 * Author: Deean
 * Date: 2023-10-14 22:50
 * FileName: LCP/LCR 024. 反转链表.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func reverseList(head *lib.ListNode) *lib.ListNode {
	type ListNode = lib.ListNode
	if head == nil || head.Next == nil {
		return head
	}
	nodes := []*ListNode{}
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	for i := len(nodes) - 1; i > 0; i-- {
		nodes[i].Next = nodes[i-1]
	}
	nodes[0].Next = nil
	return nodes[len(nodes)-1]
}

func main() {
	fmt.Println(lib.List2String(reverseList(lib.Str2List("[1,2,3,4,5]"))))
}
