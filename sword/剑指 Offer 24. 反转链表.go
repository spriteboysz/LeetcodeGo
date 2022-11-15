/**
 * Author: Deean
 * Date: 2022/11/15 22:29
 * FileName: sword/剑指 Offer 24. 反转链表.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func reverseList2(head *lib.ListNode) *lib.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func main() {
	fmt.Println(lib.List2String(reverseList2(lib.Str2List("[1,2,3,4,5]"))))
}
