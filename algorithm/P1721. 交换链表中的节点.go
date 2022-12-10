/**
 * Author: Deean
 * Date: 2022/12/10 16:35
 * FileName: algorithm/P1721. 交换链表中的节点.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func swapNodes(head *lib.ListNode, k int) *lib.ListNode {
	if head == nil {
		return nil
	}
	nodes := []*lib.ListNode{}
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	nodes[k-1], nodes[len(nodes)-k] = nodes[len(nodes)-k], nodes[k-1]
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	nodes[len(nodes)-1].Next = nil
	return nodes[0]
}

func main() {
	fmt.Println(lib.List2String(swapNodes(lib.Str2List("[7,9,6,6,7,8,3,0,9,5]"), 5)))
	fmt.Println(lib.List2String(swapNodes(lib.Str2List("[1,2]"), 1)))
}
