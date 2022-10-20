/**
 * Author: Deean
 * Date: 2022-10-20 23:23
 * FileName: algorithm/P0876. 链表的中间结点.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func middleNode(head *lib.ListNode) *lib.ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func main() {
	head := lib.Str2List("[1,2,3,4,5]")
	fmt.Println(lib.List2String(middleNode(head)))
	head = lib.Str2List("[1,2,3,4,5,6]")
	fmt.Println(lib.List2String(middleNode(head)))
}
