/**
 * Author: Deean
 * Date: 2023-10-16 23:03
 * FileName: LCP/LCR 141. 训练计划 III.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func trainningPlan(head *lib.ListNode) *lib.ListNode {
	type ListNode = lib.ListNode
	if head == nil || head.Next == nil {
		return head
	}
	nodes := []*ListNode{}
	for head != nil {
		nodes = append([]*ListNode{head}, nodes...)
		head = head.Next
	}
	for i, n := 0, len(nodes); i < n-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	nodes[len(nodes)-1].Next = nil
	return nodes[0]
}

func main() {
	fmt.Println(lib.List2String(trainningPlan(lib.Str2List("[1,2,3,4,5]"))))
}
