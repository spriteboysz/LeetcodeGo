/**
 * Author: Deean
 * Date: 2022/12/16 22:17
 * FileName: algorithm/P0024. 两两交换链表中的节点.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func swapPairs(head *lib.ListNode) *lib.ListNode {
	if head == nil {
		return head
	}
	nodes := []*lib.ListNode{}
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	for i := 0; i < len(nodes); i += 2 {
		if i+1 >= len(nodes) {
			break
		}
		nodes[i], nodes[i+1] = nodes[i+1], nodes[i]
	}
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	nodes[len(nodes)-1].Next = nil
	return nodes[0]
}

func main() {
	fmt.Println(lib.List2String(swapPairs(lib.Str2List("[1,2,3,4]"))))
	fmt.Println(lib.List2String(swapPairs(lib.Str2List("[1,2,3]"))))
	fmt.Println(lib.List2String(swapPairs(lib.Str2List("[1]"))))
}
