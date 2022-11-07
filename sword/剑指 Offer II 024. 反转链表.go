/**
 * Author: Deean
 * Date: 2022-11-07 23:43
 * FileName: sword/剑指 Offer II 024. 反转链表.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func reverseList(head *lib.ListNode) *lib.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre = reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return pre
}

func main() {
	fmt.Println(lib.List2String(reverseList(lib.Str2List("[1,2,3,4,5]"))))
}
