/**
 * Author: Deean
 * Date: 2022/12/8 23:30
 * FileName: algorithm/P0328. 奇偶链表.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func oddEvenList(head *lib.ListNode) *lib.ListNode {
	if head == nil {
		return head
	}
	even, odd := []*lib.ListNode{}, []*lib.ListNode{}
	for head != nil {
		even = append(even, head)
		head = head.Next
		if head != nil {
			odd = append(odd, head)
			head = head.Next
		}
	}
	nodes := append(even, odd...)
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	nodes[len(nodes)-1].Next = nil
	return nodes[0]
}

func main() {
	fmt.Println(lib.List2String(oddEvenList(lib.Str2List("[2,1,3,5,6,4,7]"))))
}
