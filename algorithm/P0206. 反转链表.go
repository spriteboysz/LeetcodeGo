/**
 * Author: Deean
 * Date: 2022/11/14 23:03
 * FileName: algorithm/P0206. 反转链表.go
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
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func main() {
	fmt.Println(lib.List2String(reverseList(lib.Str2List("[1,2,3,4,5]"))))
}
