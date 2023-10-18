/**
 * Author: Deean
 * Date: 2023-10-16 22:52
 * FileName: LCP/LCR 136. 删除链表的节点.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func deleteNode(head *lib.ListNode, val int) *lib.ListNode {
	type ListNode = lib.ListNode
	dummy := &ListNode{Next: head}
	prev, curr := dummy, head
	for curr != nil {
		if curr.Val == val {
			break
		}
		prev = curr
		curr = curr.Next
	}
	prev.Next = prev.Next.Next
	return dummy.Next
}
func main() {
	fmt.Println(lib.List2String(deleteNode(lib.Str2List("[4,5,1,9]"), 5)))
}
