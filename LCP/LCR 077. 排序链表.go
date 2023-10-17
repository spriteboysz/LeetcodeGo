/**
 * Author: Deean
 * Date: 2023-10-15 22:55
 * FileName: LCP/LCR 077. 排序链表.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
	"sort"
)

func sortList(head *lib.ListNode) *lib.ListNode {
	type ListNode = lib.ListNode
	if head == nil || head.Next == nil {
		return head
	}
	nodes := []*ListNode{}
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})
	for i, node := range nodes[:len(nodes)-1] {
		node.Next = nodes[i+1]
	}
	nodes[len(nodes)-1].Next = nil
	return nodes[0]
}

func main() {
	fmt.Println(lib.List2String(sortList(lib.Str2List("[4,2,1,3]"))))
}
