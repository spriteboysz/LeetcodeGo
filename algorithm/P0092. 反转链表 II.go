/**
 * Author: Deean
 * Date: 2022/12/12 23:38
 * FileName: algorithm/P0092. 反转链表 II.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func reverseBetween(head *lib.ListNode, left int, right int) *lib.ListNode {
	nodes := []*lib.ListNode{}
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	left, right = left-1, right-1
	for i, j := left, right; i <= left+(right-left)/2; i, j = i+1, j-1 {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	}

	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	nodes[len(nodes)-1].Next = nil
	return nodes[0]
}

func main() {
	fmt.Println(lib.List2String(reverseBetween(lib.Str2List("[1,2,3,4,5]"), 2, 4)))
}
