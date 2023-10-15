/**
 * Author: Deean
 * Date: 2023-10-14 22:27
 * FileName: LCP/LCR 021. 删除链表的倒数第 N 个结点.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func removeNthFromEnd(head *lib.ListNode, n int) *lib.ListNode {
	dummy := &lib.ListNode{Val: -1, Next: head}
	fast, slow := dummy, dummy
	for n > 0 {
		fast = fast.Next
		n--
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func main() {
	fmt.Println(lib.List2String(removeNthFromEnd(lib.Str2List("[1,2,3,4,5]"), 2)))
}
