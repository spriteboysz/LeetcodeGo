/**
 * Author: Deean
 * Date: 2023-10-15 23:06
 * FileName: LCP/LCR 078. 合并 K 个升序链表.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
	"sort"
)

func mergeKLists(lists []*lib.ListNode) *lib.ListNode {
	type ListNode = lib.ListNode
	if len(lists) == 0 {
		return nil
	}
	nodes := []*ListNode{}
	for _, list := range lists {
		cur := list
		for cur != nil {
			nodes = append(nodes, cur)
			cur = cur.Next
		}
	}
	if len(nodes) == 0 {
		return nil
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
	l1 := lib.Str2List("[1,4,5]")
	l2 := lib.Str2List("[1,3,4]")
	l3 := lib.Str2List("[2,6]")
	fmt.Println(lib.List2String(mergeKLists([]*lib.ListNode{l1, l2, l3})))
}
