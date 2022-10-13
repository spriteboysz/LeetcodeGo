/**
 * Author: Deean
 * Date: 2022-10-13 00:01
 * FileName: interview/面试题 02.03. 删除中间节点.go
 * Description:
 */

package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	fmt.Println()
}
