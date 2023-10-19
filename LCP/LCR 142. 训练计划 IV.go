/**
 * Author: Deean
 * Date: 2023-10-16 23:10
 * FileName: LCP/LCR 142. 训练计划 IV.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
	"sort"
)

func trainningPlan2(l1 *lib.ListNode, l2 *lib.ListNode) *lib.ListNode {
	type ListNode = lib.ListNode
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	nodes := []*ListNode{}
	for l1 != nil {
		nodes = append(nodes, l1)
		l1 = l1.Next
	}
	for l2 != nil {
		nodes = append(nodes, l2)
		l2 = l2.Next
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})
	for i, n := 0, len(nodes); i < n-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	nodes[len(nodes)-1].Next = nil
	return nodes[0]
}

func main() {
	l1, l2 := lib.Str2List("[1,2,4]"), lib.Str2List("[1,3,4]")
	fmt.Println(lib.List2String(trainningPlan2(l1, l2)))
}
