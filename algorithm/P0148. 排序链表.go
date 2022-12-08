/**
 * Author: Deean
 * Date: 2022/12/8 23:00
 * FileName: algorithm/P0148. 排序链表.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
	"sort"
)

func sortList(head *lib.ListNode) *lib.ListNode {
	if head == nil {
		return head
	}
	nodes := []*lib.ListNode{}
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	nodes[len(nodes)-1].Next = nil
	return nodes[0]
}

func main() {
	fmt.Println(lib.List2String(sortList(lib.Str2List("[4,2,1,3]"))))
}
