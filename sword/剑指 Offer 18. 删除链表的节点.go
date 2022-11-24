/**
 * Author: Deean
 * Date: 2022/11/24 22:57
 * FileName: sword/剑指 Offer 18. 删除链表的节点.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func deleteNode(head *lib.ListNode, val int) *lib.ListNode {
	dummy := &lib.ListNode{Next: head}
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
