/**
 * Author: Deean
 * Date: 2022-10-14 23:49
 * FileName: algorithm/P0237. 删除链表中的节点.go
 * Description:
 */

package main

import (
	"leetcode/lib"
)

type ListNode = lib.ListNode

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	node := &ListNode{}
	deleteNode(node)
}
