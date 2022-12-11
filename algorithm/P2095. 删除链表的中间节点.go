/**
 * Author: Deean
 * Date: 2022/12/11 17:51
 * FileName: algorithm/P2095. 删除链表的中间节点.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func deleteMiddle(head *lib.ListNode) *lib.ListNode {
	dummy := &lib.ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func main() {
	fmt.Println(lib.List2String(deleteMiddle(lib.Str2List("[1,3,4,7,1,2,6]"))))
}
